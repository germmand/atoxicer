package handlers

import (
	"context"
	"log"
	"os"

	"github.com/germmand/atoxicer/bot/helpers"

	"github.com/bwmarrin/discordgo"
	"github.com/germmand/atoxicer/bot/constants"
	"github.com/germmand/atoxicer/bot/embeds"
	"github.com/germmand/atoxicer/firebase"
	"github.com/germmand/atoxicer/perspective"
)

// MessageCreateHandler handles all of the logic for muting a user
// if being too toxic.
func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	perspectiveKey := os.Getenv("PERSPECTIVE_API_KEY")
	perspectiveSession := perspective.New(perspectiveKey)
	toxicity, err := perspectiveSession.ObtainToxicity(m.Content)
	if err != nil {
		return
	}

	toxicityLevel := toxicity.AttributeScores.Toxicity.SummaryScore.Value
	toxicityType := constants.DetermineToxicType(toxicityLevel)
	if toxicityType == constants.NonToxic {
		return
	}
	embedConfig := constants.EmbedConfigTypes[toxicityType]

	ctx := context.Background()
	firebaseApp := firebase.NewApp(ctx)
	firestoreApp := firebaseApp.NewFirestoreSession(ctx)

	defer firestoreApp.FirestoreSession.Close()

	warningUser, err := firestoreApp.RetrieveWarning(ctx, m.Author.ID, m.GuildID)
	if err != nil {
		log.Printf("An error has occurred reading data: %s", err)
	}

	// We put this here because if UserID is "" it means that the record
	// does not exist in Firestore.
	if warningUser.UserID == "" {
		warningUser.UserID = m.Author.ID
		warningUser.GuildID = m.GuildID
	}

	warningUser.UpdateWarningUponToxicity(toxicityType)
	var embedMessageConfig embeds.Embed

	if warningUser.RedWarnings >= 3 {
		err = firestoreApp.DeleteWarning(ctx, m.Author.ID)
		if err != nil {
			log.Fatalln(err)
			return
		}

		roles, err := s.GuildRoles(m.GuildID)
		if err != nil {
			log.Fatalln(err)
		}

		mutedRoleID, err := helpers.FilterRoleByName(roles, os.Getenv("DISCORD_MUTED_ROLE"))
		if err != nil {
			log.Fatalln(err)
			return
		}

		err = s.GuildMemberRoleAdd(warningUser.GuildID, warningUser.UserID, mutedRoleID)
		if err != nil {
			log.Fatalln(err)
			return
		}

		embedMessageConfig = &embeds.MutedEmbedConfig{
			MessageData: m,
		}
	} else {
		err = firestoreApp.SetWarning(ctx, m.Author.ID, warningUser)
		if err != nil {
			log.Printf("An error has occurred updating warning: %s", err)
		}

		embedMessageConfig = &embeds.WarningEmbedConfig{
			MessageData:        m,
			EmbedGeneralConfig: embedConfig,
			WarningModel:       warningUser,
			ToxicityScore:      toxicityLevel,
		}
	}

	embedMessage := embedMessageConfig.GenerateEmbed()

	messageSend := &discordgo.MessageSend{
		Embed: embedMessage,
		AllowedMentions: &discordgo.MessageAllowedMentions{
			Parse: []discordgo.AllowedMentionType{discordgo.AllowedMentionTypeUsers},
		},
	}

	_, err = s.ChannelMessageSendComplex(m.ChannelID, messageSend)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

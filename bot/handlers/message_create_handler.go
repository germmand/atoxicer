package handlers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/germmand/atoxicer/bot/constants"
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

	embedMessage := &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       "Advertencia",
		Description: fmt.Sprintf("<@%s>, tu mensaje fue detectado como toxico.", m.Author.ID),
		Color:       embedConfig.Color,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "Mensaje",
				Value: fmt.Sprintf("```\n%s```\n", m.Content),
			},
			&discordgo.MessageEmbedField{
				Name:   "Porcentaje de toxicidad",
				Value:  fmt.Sprintf("%.2f%%", toxicityLevel*100),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Toxicidad",
				Value:  embedConfig.ToxicityLevel,
				Inline: true,
			},
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name: "Atoxicer",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Advertencia 1/3",
		},
	}

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

	err = firestoreApp.SetWarning(ctx, m.Author.ID, warningUser)
	if err != nil {
		log.Printf("An error has occurred updating warning: %s", err)
	}

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

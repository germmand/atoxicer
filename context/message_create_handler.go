package context

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/germmand/atoxicer/bot/constants"
)

func (c *AtoxicerContext) MessageCreate() func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		toxicity, err := c.PerspectiveSession.ObtainToxicity(m.Content)
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

		messageSend := &discordgo.MessageSend{
			Embed: embedMessage,
			AllowedMentions: &discordgo.MessageAllowedMentions{
				Parse: []discordgo.AllowedMentionType{discordgo.AllowedMentionTypeUsers},
			},
		}

		message, err := s.ChannelMessageSendComplex(m.ChannelID, messageSend)
		if err != nil {
			log.Fatalln(err)
			return
		}

		fmt.Println("Success", message)
	}
}

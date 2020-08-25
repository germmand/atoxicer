package context

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
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

		// TODO: Separar toda la logica de discordgo a un paquete (bot) separado...

		embedMessage := &discordgo.MessageEmbed{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Advertencia",
			Description: fmt.Sprintf("<@%s>, tu mensaje fue detectado como toxico.", m.Author.ID),
			// Color:       10878976, Rojo
			// Color: 12893718, Amarillo
			Color: 1491996,
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
					Name:   "Nivel",
					Value:  "Leve",
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

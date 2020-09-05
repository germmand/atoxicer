package embeds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/germmand/atoxicer/bot/constants"
	"github.com/germmand/atoxicer/firebase/firestore/models"
)

type WarningEmbedConfig struct {
	MessageData        *discordgo.MessageCreate
	EmbedGeneralConfig *constants.EmbedConfig
	WarningModel       *models.Warning
	ToxicityScore      float64
}

func (w *WarningEmbedConfig) GenerateEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       "Advertencia",
		Description: fmt.Sprintf("<@%s>, tu mensaje fue detectado como toxico.", w.MessageData.Author.ID),
		Color:       w.EmbedGeneralConfig.Color,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "Mensaje",
				Value: fmt.Sprintf("```\n%s```\n", w.MessageData.Content),
			},
			&discordgo.MessageEmbedField{
				Name:   "Porcentaje de toxicidad",
				Value:  fmt.Sprintf("%.2f%%", w.ToxicityScore*100),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Toxicidad",
				Value:  w.EmbedGeneralConfig.ToxicityLevel,
				Inline: true,
			},
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name: "Atoxicer",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Infracciones %d/3 - Advertencias %d/2", w.WarningModel.RedWarnings, w.WarningModel.YellowWarnings),
		},
	}
}

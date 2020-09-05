package embeds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type MutedEmbedConfig struct {
	MessageData *discordgo.MessageCreate
}

func (m *MutedEmbedConfig) GenerateEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       "Muteado",
		Description: fmt.Sprintf("<@%s>, has sido penalizado debido a tu comportamiento.", m.MessageData.Author.ID),
		Color:       14903819, // Orange
		Image: &discordgo.MessageEmbedImage{
			URL: "https://media.giphy.com/media/LMi8xcUpr0cg4XvNUu/giphy.gif",
		},
	}
}

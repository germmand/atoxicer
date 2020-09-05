package embeds

import (
	"github.com/bwmarrin/discordgo"
)

type Embed interface {
	GenerateEmbed() *discordgo.MessageEmbed
}

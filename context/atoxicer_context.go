package context

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/germmand/atoxicer/bot"
	"github.com/germmand/atoxicer/perspective"
)

type AtoxicerContext struct {
	BotSession         *discordgo.Session
	PerspectiveSession *perspective.PerspectiveSession
}

func New() *AtoxicerContext {
	perspectiveKey := os.Getenv("PERSPECTIVE_API_KEY")
	perspectiveSession := perspective.New(perspectiveKey)

	token := os.Getenv("DISCORD_TOKEN_BOT")
	botSession := bot.New(token)

	atoxicerContext := &AtoxicerContext{
		PerspectiveSession: perspectiveSession,
		BotSession:         botSession,
	}

	return atoxicerContext
}

func (c *AtoxicerContext) SetupHandlers() {
	c.BotSession.AddHandler(c.MessageCreate())

	c.BotSession.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
}

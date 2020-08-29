package bot

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/germmand/atoxicer/bot/handlers"
)

func New(token string) *discordgo.Session {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating discord session", err)
		os.Exit(1)
	}

	dg.AddHandler(handlers.MessageCreateHandler)
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	return dg
}

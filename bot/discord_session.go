package bot

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func New(token string) *discordgo.Session {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating discord session", err)
		os.Exit(1)
	}
	return dg
}

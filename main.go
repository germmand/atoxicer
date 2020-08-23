package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/germmand/atoxicer/perspective"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN_BOT")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating discord session", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection", err)
		return
	}

	fmt.Println("Bot is now running. Press CMD+C to Exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	perspectiveKey := os.Getenv("PERSPECTIVE_API_KEY")
	perspectiveSession := perspective.New(perspectiveKey)

	if m.Author.ID == s.State.User.ID {
		return
	}

	toxicity, err := perspectiveSession.ObtainToxicity(m.Content)
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
				Value: m.Content,
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

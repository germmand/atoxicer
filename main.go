package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/bwmarrin/discordgo"
    "github.com/germmand/atoxicer/perspective"
)

func main () {
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

    // Esto obviamente tiene que refactorizarse a una estructura aparte...
    toxicityLevel := toxicity["attributeScores"].(map[string]interface{})["TOXICITY"].(map[string]interface{})["summaryScore"].(map[string]interface{})["value"]
    s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Tu mensaje tiene un porcentaje de toxicidad de: %.2f%%", toxicityLevel.(float64) * 100))
}

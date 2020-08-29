package context

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/bwmarrin/discordgo"
	"github.com/germmand/atoxicer/bot"
	"github.com/germmand/atoxicer/perspective"
)

type AtoxicerContext struct {
	BotSession         *discordgo.Session
	PerspectiveSession *perspective.PerspectiveSession
	FirestoreSession   *firestore.Client
}

func New() *AtoxicerContext {
	perspectiveKey := os.Getenv("PERSPECTIVE_API_KEY")
	perspectiveSession := perspective.New(perspectiveKey)

	token := os.Getenv("DISCORD_TOKEN_BOT")
	botSession := bot.New(token)

	// Separate this into its own package (probably)
	ctx := context.Background()
	firebaseApp, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	firestore, err := firebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	/*
		// This does work here ... :/
		_, _, err = firestore.Collection("users").Add(ctx, map[string]interface{}{
			"first": "Ada",
			"last":  "Lovelace",
			"born":  1815,
		})
		if err != nil {
			log.Fatalf("Failed adding alovelace: %v", err)
		}
	*/

	atoxicerContext := &AtoxicerContext{
		PerspectiveSession: perspectiveSession,
		BotSession:         botSession,
		FirestoreSession:   firestore,
	}

	return atoxicerContext
}

func (c *AtoxicerContext) SetupHandlers() {
	c.BotSession.AddHandler(c.MessageCreate())

	c.BotSession.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
}

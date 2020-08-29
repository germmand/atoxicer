package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/germmand/atoxicer/bot"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN_BOT")
	dg := bot.New(token)

	err := dg.Open()
	if err != nil {
		fmt.Println("error opening connection", err)
		return
	}

	defer dg.Close()

	fmt.Println("Bot is now running. Press CMD+C to Exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

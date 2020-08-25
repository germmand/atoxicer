package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/germmand/atoxicer/context"
)

var Context *context.AtoxicerContext

func init() {
	Context = context.New()
	Context.SetupHandlers()
}

func main() {
	dg := Context.BotSession

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

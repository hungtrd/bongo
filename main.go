package main

import (
	"bongo/config"
	"bongo/pkg/discord"
	"bongo/pkg/openai"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.Setup()
	openai.Setup()
	dg := discord.Setup()
	// Open a websocket connection to Discord and begin listening.
	err := dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Run cronjob
	c := cronjobSetup()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	defer dg.Close()

	defer c.Stop()
}

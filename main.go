package main

import (
	"discord-bot/pkg/simsimi"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var token string

func main() {

	token = ""
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.Identify.Intents |= discordgo.IntentMessageContent

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	fmt.Println(m.Author, ": ", m.Content)
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check message send to BOT
	if !strings.HasPrefix(m.Content, "!bot") {
		return
	}

	_, content, ok := strings.Cut(m.Content, "!bot")
	if !ok {
		return
	}
	content = strings.TrimSpace(content)

	// If the message is "ping" reply with "Pong!"
	if content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
		return
	}

	// If the message is "pong" reply with "Ping!"
	if content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
		return
	}

	msg, err := simsimi.SendMessage(content)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	s.ChannelMessageSend(m.ChannelID, msg)
}

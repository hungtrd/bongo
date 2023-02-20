package discord

import (
	"bongo/config"
	"bongo/pkg/simsimi"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var Session *discordgo.Session

func Setup() *discordgo.Session {

	dg, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return nil
	}

	// Export discord session
	Session = dg

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.Identify.Intents |= discordgo.IntentMessageContent

	return dg
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

	msg, err := simsimi.SendMessage(content)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	s.ChannelMessageSend(m.ChannelID, msg)
}

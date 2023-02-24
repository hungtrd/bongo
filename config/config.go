package config

import "os"

const (
	DiscordRemindChannelID string = "1070606289048186950"
)

var (
	DiscordToken string
	OpenAIToken  string
)

func getConfig() {
	DiscordToken = os.Getenv("DISCORD_TOKEN")
	OpenAIToken = os.Getenv("OPENAI_TOKEN")

	// Get int
	// var err error
}

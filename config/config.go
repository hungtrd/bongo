package config

import "os"

const (
	DiscordRemindChannelID string = "1098535545933672499"
	TimeZone               string = "Asia/Ho_Chi_Minh"
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

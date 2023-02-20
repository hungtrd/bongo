package main

import (
	"bongo/config"
	"bongo/pkg/discord"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

func cronjobSetup() *cron.Cron {
	c := cron.New(cron.WithSeconds())

	jobID, _ := c.AddFunc("0 * * * * *", func() {
		fmt.Println("Every minute", time.Now())
		JobLunarCalendarCheck()
	})
	fmt.Println("Sign job: ", jobID)
	jobID, _ = c.AddFunc("0 0 8 * * *", func() { JobLunarCalendarCheck() })
	fmt.Println("Sign job: ", jobID)

	c.Start()

	return c
}

func JobLunarCalendarCheck() {
	users := []string{"394490825960325130"}
	roles := []string{"1077189580425527306"}

	dg := discord.Session
	msg := fmt.Sprintf("<@&%s> Hello! It time is %v", roles[0], time.Now())

	data := discordgo.MessageSend{
		Content: msg,
		AllowedMentions: &discordgo.MessageAllowedMentions{
			Users:       users,
			Roles:       roles,
			RepliedUser: true,
		},
	}
	_, err := dg.ChannelMessageSendComplex(config.DiscordRemindChannelID, &data)
	if err != nil {
		fmt.Println("[main.JobLunarCalendarCheck] failed: ", err)
	}
}

package main

import (
	"fmt"
	"time"

	"bongo/config"
	"bongo/pkg/discord"

	"github.com/bwmarrin/discordgo"
	"github.com/hungtrd/amlich"
	"github.com/robfig/cron/v3"
)

var timeLoc *time.Location

func cronjobSetup() *cron.Cron {
	var err error

	timeLoc, err = time.LoadLocation(config.TimeZone)
	if err != nil {
		panic(err)
	}

	c := cron.New(cron.WithSeconds(), cron.WithLocation(timeLoc))

	jobID, _ := c.AddFunc("0 0 8 * * *", func() { JobLunarCalendarCheck() })
	fmt.Println("Sign job: ", jobID)

	c.Start()

	return c
}

func JobLunarCalendarCheck() {
	users := []string{"394490825960325130"}
	roles := []string{"1077189580425527306"}
	send := false
	msg := ""

	// check today
	today := time.Now().In(timeLoc)
	todayLunarD, todayLunarM, todayLunarY, _ := amlich.Solar2Lunar(today.Day(), int(today.Month()), today.Year(), 7)
	if todayLunarD == 15 || todayLunarD == 1 {
		send = true
		msg = fmt.Sprintf("<@&%s> Hôm nay âm lịch là ngày %v tháng %v năm %v. Nhớ mua đồ thắp hương!", roles[0], todayLunarD, todayLunarM, todayLunarY)
	}

	// check tomorrow
	tomor := time.Now().AddDate(0, 0, 1).In(timeLoc)
	tomorLunarD, tomorLunarM, tomorLunarY, _ := amlich.Solar2Lunar(tomor.Day(), int(tomor.Month()), tomor.Year(), 7)
	if tomorLunarD == 15 || tomorLunarD == 1 {
		send = true
		msg = fmt.Sprintf("<@&%s> Ngày mai âm lịch là ngày %v tháng %v năm %v. Nhớ mua đồ thắp hương!", roles[0], tomorLunarD, tomorLunarM, tomorLunarY)
	}

	if send {
		dg := discord.Session

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
}

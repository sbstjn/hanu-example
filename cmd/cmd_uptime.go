package cmd

import (
	"time"

	"github.com/sbstjn/hanu/bot"
)

func init() {
	Register(
		"uptime",
		"Reply with the uptime",
		func(conv *bot.Conversation) {
			conv.Reply("Thanks for asking! I'm running since `%s`", time.Since(Start))
		},
	)
}

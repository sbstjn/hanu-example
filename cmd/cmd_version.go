package cmd

import "github.com/sbstjn/hanu/bot"

func init() {
	Register(
		"version",
		"Reply with the current chatbot version",
		func(conv *bot.Conversation) {
			conv.Reply("Thanks for asking! I'm running with `%s`", Version)
		},
	)
}

package cmd

import "github.com/sbstjn/hanu/bot"

func init() {
	Register(
		"hi",
		"Reply with a lovely welcome greeting",
		func(conv bot.Conversation) {
			conv.Reply("Hi yourself!")
		},
	)
}

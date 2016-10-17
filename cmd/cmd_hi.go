package cmd

import "github.com/sbstjn/hanu"

func init() {
	Register(
		"hi",
		"Reply with a lovely welcome greeting",
		func(conv hanu.ConversationInterface) {
			conv.Reply("Hi yourself!")
		},
	)
}

package cmd

import "github.com/sbstjn/hanu/conversation"

func init() {
	Register(
		"hi",
		"Reply with a lovely welcome greeting",
		func(conv conversation.Interface) {
			conv.Reply("Hi yourself!")
		},
	)
}

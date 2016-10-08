package cmd

import "github.com/sbstjn/hanu/conversation"

func init() {
	Register(
		"version",
		"Reply with the current chatbot version",
		func(conv conversation.Interface) {
			conv.Reply("Thanks for asking! I'm running with `%s`", Version)
		},
	)
}

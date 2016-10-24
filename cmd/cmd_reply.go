package cmd

import "github.com/sbstjn/hanu"

func init() {
	Register(
		"reply (.*)",
		"Reply with the passed message",
		func(conv hanu.ConversationInterface) {
			data, _ := conv.Match(0)
			conv.Reply(data)
		},
	)
}

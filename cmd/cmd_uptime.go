package cmd

import (
	"time"

	"github.com/sbstjn/hanu"
)

func init() {
	Register(
		"uptime",
		"Reply with the uptime",
		func(conv hanu.ConversationInterface) {
			conv.Reply("Thanks for asking! I'm running since `%s`", time.Since(Start))
		},
	)
}

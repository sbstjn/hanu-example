package cmd

import (
	"time"

	"github.com/sbstjn/hanu/conversation"
)

func init() {
	Register(
		"uptime",
		"Reply with the uptime",
		func(conv conversation.Interface) {
			conv.Reply("Thanks for asking! I'm running since `%s`", time.Since(Start))
		},
	)
}

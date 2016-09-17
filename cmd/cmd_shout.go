package cmd

import (
	"strings"

	"github.com/sbstjn/hanu/bot"
)

func init() {
	Register(
		"shout <word>",
		"Reply the password in uppercase letters",
		func(conv *bot.Conversation) {
			conv.Reply(strings.ToUpper(conv.Param("word")))
		},
	)
}

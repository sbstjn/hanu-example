package cmd

import (
	"strings"

	"github.com/sbstjn/hanu/bot"
)

func init() {
	Register(
		"whisper <word>",
		"Reply the passed word in lowercase letters",
		func(conv *bot.Conversation) {
			conv.Reply(strings.ToLower(conv.Param("word")))
		},
	)
}

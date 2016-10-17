package cmd

import (
	"strings"

	"github.com/sbstjn/hanu"
)

func init() {
	Register(
		"whisper <word>",
		"Reply the passed word in lowercase letters",
		func(conv hanu.ConversationInterface) {
			str, _ := conv.String("word")
			conv.Reply(strings.ToLower(str))
		},
	)
}

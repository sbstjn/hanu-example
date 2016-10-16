package cmd

import (
	"strings"

	"github.com/sbstjn/hanu/conversation"
)

func init() {
	Register(
		"whisper <word>",
		"Reply the passed word in lowercase letters",
		func(conv conversation.Interface) {
			str, _ := conv.String("word")
			conv.Reply(strings.ToLower(str))
		},
	)
}

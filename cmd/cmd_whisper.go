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
			conv.Reply(strings.ToLower(conv.Param("word")))
		},
	)
}

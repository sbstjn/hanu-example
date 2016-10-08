package cmd

import (
	"strings"

	"github.com/sbstjn/hanu/conversation"
)

func init() {
	Register(
		"shout <word>",
		"Reply the password in uppercase letters",
		func(conv conversation.Interface) {
			conv.Reply(strings.ToUpper(conv.Param("word")))
		},
	)
}

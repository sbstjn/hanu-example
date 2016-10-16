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
			str, _ := conv.String("word")
			conv.Reply(strings.ToUpper(str))
		},
	)
}

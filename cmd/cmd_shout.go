package cmd

import (
	"strings"

	"github.com/sbstjn/hanu"
)

func init() {
	Register(
		"shout <word>",
		"Reply the password in uppercase letters",
		func(conv hanu.ConversationInterface) {
			str, _ := conv.String("word")
			conv.Reply(strings.ToUpper(str))
		},
	)
}

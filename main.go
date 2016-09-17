package main

import (
	"log"

	"github.com/sbstjn/hanu"
	"github.com/sbstjn/hanu-example/cmd"
)

var (
	// Version is the bot version
	Version = "0.0.1"
)

func main() {
	bot, err := hanu.New("SLACK_API_TOKEN")

	if err != nil {
		log.Fatal(err)
	}

	cmd.Version = Version
	cmdList := cmd.List()
	for i := 0; i < len(cmdList); i++ {
		bot.Register(cmdList[i])
	}

	bot.Listen()
}

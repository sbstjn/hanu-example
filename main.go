package main

import (
	"log"
	"time"

	"github.com/sbstjn/hanu"
	"github.com/sbstjn/hanu-example/cmd"
	"github.com/spf13/viper"
)

var (
	// Version is the bot version
	Version = "0.0.1"
	// SlackToken will be set by ENV or config file in init()
	SlackToken = ""
)

func init() {
	viper.SetConfigName(".hanu-example")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	SlackToken = viper.GetString("HANU_EXAMPLE_SLACK_TOKEN")
}

func main() {
	bot, err := hanu.New(SlackToken)

	if err != nil {
		log.Fatal(err)
	}

	cmd.Version = Version
	cmd.Start = time.Now()
	cmdList := cmd.List()
	for i := 0; i < len(cmdList); i++ {
		bot.Register(cmdList[i])
	}

	bot.Listen()
}

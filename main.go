package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

var (
	port string
)

func startWeb() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK.")
	})

	r.Run(":" + port)
}

func startBot() {
	bot, err := hanu.New(SlackToken)

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

func init() {
	port = os.Getenv("PORT")
}

func main() {
	startWeb()
	startBot()
}

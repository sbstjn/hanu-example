package cmd

import (
	"time"

	"github.com/sbstjn/hanu/bot"
	"github.com/sbstjn/platzhalter"
)

var commandList []bot.Command

// Version stores the chatbot version. Will be updated by `main.go`
var Version string

// Start contains the Time when the process started
var Start time.Time

// Register adds a new command to commandList
func Register(command string, description string, handler bot.CommandHandler) {
	commandList = append(commandList, bot.Command{
		Command:     platzhalter.NewCommand(command),
		Description: description,
		Handler:     handler,
	})
}

// List returns commandList
func List() []bot.Command {
	return commandList
}

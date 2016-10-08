package cmd

import (
	"time"

	"github.com/sbstjn/hanu/command"
)

var commandList []command.Interface

// Version stores the chatbot version. Will be updated by `main.go`
var Version string

// Start contains the Time when the process started
var Start time.Time

// Register adds a new command to commandList
func Register(cmd string, description string, handler command.Handler) {
	commandList = append(commandList, command.New(cmd, description, handler))
}

// List returns commandList
func List() []command.Interface {
	return commandList
}

package cmd

import (
	"time"

	"github.com/sbstjn/hanu"
)

var commandList []hanu.CommandInterface

// Version stores the chatbot version. Will be updated by `main.go`
var Version string

// Start contains the Time when the process started
var Start time.Time

// Register adds a new command to commandList
func Register(cmd string, description string, handler hanu.Handler) {
	commandList = append(commandList, hanu.NewCommand(cmd, description, handler))
}

// List returns commandList
func List() []hanu.CommandInterface {
	return commandList
}

package platzhalter

// NewCommand creates a new Command
func NewCommand(text string) Command {
	return Command{
		Text: text,
	}
}

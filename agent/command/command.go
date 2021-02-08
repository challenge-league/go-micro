// Package command is an interface for defining bot commands
package command

import (
	nakama "github.com/challenge-league/nakama-go/context"
)

var (
	// Commmands keyed by golang/regexp patterns
	// regexp.Match(key, input) is used to match
	Commands = map[string]Command{}
)

// Command is the interface for specific named
// commands executed via plugins or the bot.
type Command interface {
	// Executes the command with args passed in
	Exec(nakamaCtx *nakama.Context, args ...string) ([]byte, error)
	// Usage of the command
	Usage() string
	// Description of the command
	Description() string
	// Name of the command
	String() string
}

type cmd struct {
	name        string
	usage       string
	description string
	exec        func(nakamaCtx *nakama.Context, args ...string) ([]byte, error)
}

func (c *cmd) Description() string {
	return c.description
}

func (c *cmd) Exec(nakamaCtx *nakama.Context, args ...string) ([]byte, error) {
	return c.exec(nakamaCtx, args...)
}

func (c *cmd) Usage() string {
	return c.usage
}

func (c *cmd) String() string {
	return c.name
}

// NewCommand helps quickly create a new command
func NewCommand(name, usage, description string, exec func(nakamaCtx *nakama.Context, args ...string) ([]byte, error)) Command {
	return &cmd{
		name:        name,
		usage:       usage,
		description: description,
		exec:        exec,
	}
}

// Package commands base command features
// GetName, Clear, Parse, Verify, Run
package commands

import (
	"msg"
	"strings"
)

// ICommand for base command's required behaviour
type ICommand interface {
	Help()
	GetName() string
	Parse(string) error
	Verify() error
	Run() (string, error)
	Clear()
	isHelp() bool
}

// Command object structure
type Command struct {
	Cmd,
	InputArgs string
	OutPut string
	Args   []string
}

// NewCommand to create command instance
func NewCommand() *Command {
	var cmd = new(Command)
	return cmd
}

// Help to show usage
func (c *Command) Help() {
}

// GetName to get the command name
func (c *Command) GetName() string {
	return c.Cmd
}

// Clear to clear the history data
func (c *Command) Clear() {
	c.InputArgs = msg.Empty
	c.Args = []string{}
	c.OutPut = msg.Empty
}

// Parse to help command to parse arguments from input string
func (c *Command) Parse(argString string) error {
	c.InputArgs = argString
	c.Args = strings.Split(argString, msg.Seperator)
	for index, arg := range c.Args {
		c.Args[index] = strings.TrimSpace(arg)
	}
	return nil
}

// Verify the provided Arguments
func (c *Command) Verify() error {
	return nil
}

// Run the command with arguments
func (c *Command) Run() (string, error) {
	return c.OutPut, nil
}

// check and show command help
func (c *Command) isHelp() bool {
	if (strings.ToLower(c.Args[0]) == "help") {
		c.Help()
		return true;
	}
	return false
}
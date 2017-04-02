// Package commands implements the basic shell command execution process
// Includes help, parse, verify, run
package commands

import (
	"bufio"
	"msg"
	"strings"
)

// IManager should have behaviour run and a base parse
type IManager interface {
	Parse() error
	Run(string, *bufio.Scanner) (string, error)
}

// Manager handles requested command and available command's list
type Manager struct {
	cmd, argString string
	Commands       map[string]ICommand
}

var mgrCmd = &Manager{
	Commands: make(map[string]ICommand),
}

// NewManager return command manager
func NewManager() *Manager {
	mgrCmd.Register(NewCmdGetPermission())
	mgrCmd.Register(NewCmdSetPermission())
	mgrCmd.Register(NewCmdHelp())
	return mgrCmd
}

// Register Command registration with manager
func (cm *Manager) Register(cmd ICommand) {
	cmdName := strings.ToLower(cmd.GetName())
	cm.Commands[cmdName] = cmd
}

// IsValidCommad verifies the requested command is valid or not
func (cm *Manager) IsValidCommad(cmdName string) bool {
	cmdName = strings.ToLower(cmdName)
	_, ok := cm.Commands[cmdName]
	return ok
}

// Parse requested command and arguments
func (cm *Manager) Parse(cmdString string) error {
	cmdString = strings.Trim(cmdString, " \n\t")
	results := strings.SplitN(cmdString, msg.Space, 2)
	cm.cmd = strings.ToLower(results[0])
	if len(results) > 1 {
		cm.argString = results[1]
	}
	if msg.Empty == cm.cmd {
		return msg.ErrInvalidCommand
	}
	return nil
}

// Run the requested command and provide output
func (cm *Manager) Run(cmdString string, reader *bufio.Scanner) (string, error) {
	err := cm.Parse(cmdString)
	if nil != err {
		return msg.Empty, err
	}
	cmd, ok := cm.Commands[cm.cmd]
	if ok {
		cmd.Clear()
		err := cmd.Parse(cm.argString, reader)
		if nil != err {
			return msg.Empty, msg.ErrInvalidParams
		}
		if nil == cmd.Verify() {
			return cmd.Run()
		}
		return msg.Empty, msg.ErrInvalidParams
	}
	return msg.Empty, msg.ErrInvalidCommand
}

// Package commands with 'get permission' command implementation
package commands

import (
	"bufio"
	"fmt"
	"msg"
)

// CmdGetPermission defined arguments and related methods
type CmdGetPermission struct {
	Command
}

// NewCmdGetPermission new get permission command instance
func NewCmdGetPermission() *CmdGetPermission {
	var cmd = new(CmdGetPermission)
	cmd.Cmd = "permissions_of"
	return cmd
}

// Help to print help of get permission command
func (cp *CmdGetPermission) Help() {
	fmt.Println(`🔸  permissions_of -
	
  This allows to set permissions for distributor.
  eg: Permissions_of DISTRIBUTOR1
  This allows to get all permissions of DISTRIBUTOR1.
`)
}

// Parse to parse arguments
func (cp *CmdGetPermission) Parse(argString string, reader *bufio.Scanner) error {
	cp.Command.Parse(argString, reader)
	return nil
}

// Verify to check the provided parameters are valid or not
func (cp *CmdGetPermission) Verify() error {
	if 1 != len(cp.Args) {
		return msg.ErrInvalidParams
	}
	return nil
}

// Run to execute the command and provide result
func (cp *CmdGetPermission) Run() (string, error) {
	if cp.isHelp() {
		cp.Help()
		return "", nil
	}
	return "No output", nil
}

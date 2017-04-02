// Package commands with 'get permission' command implementation
package commands

import (

)
import (
	"fmt"
)

// Help defined arguments and related methods
type CmdHelp struct {
	Command
}

// NewCmdGetPermission new park command instance
func NewCmdHelp() *CmdHelp {
	var cmd = new(CmdHelp)
	cmd.Cmd = "help"
	return cmd
}

// Help to print help of get permission command
func (cp *CmdHelp) Help() {
	fmt.Print(`🔸  help -
	Shows command help
`)
}

// Parse to parse arguments
func (cp *CmdHelp) Parse(argString string) error {
	cp.Command.Parse(argString)
	return nil
}

// Verify to check the provided parameters are valid or not
func (cp *CmdHelp) Verify() error {
	return nil;
}

// Run to execute the command and provide result
func (cp *CmdHelp) Run() (string, error) {
	for _, cmd := range mgrCmd.Commands {
		cmd.Help()
	}
	return "", nil
}

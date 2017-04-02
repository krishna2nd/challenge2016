// Package commands with shell implementations
package commands

import (
	"bufio"
	"fmt"
	"os"
	"msg"
	"strings"
)

// Shell stores current shell instance information
type Shell struct {
	PS1 string
	// Can be used for command history
	// cmdList []string
}

// NewShell create a shell
func NewShell() *Shell {
	return &Shell{
		PS1: ">",
		// cmdList: make([]string, 0),
	}
}

// Process method to handle commands
func (sh *Shell) Process() error {
	reader := bufio.NewReader(os.Stdin)
	cmdMgr := NewManager()
	sh.prompt()
	for {
		cmdInput, _ := reader.ReadString('\n')
		cmdInput = strings.TrimRight(cmdInput, msg.NewLine)
		if msg.Empty != cmdInput {
		} else {
			cmdInput = "help"
		}
		out, err := cmdMgr.Run(cmdInput)
		processOutput(out, err)
		sh.prompt()
	}
	return nil
}

// prompt display command prompt PS1
func (sh *Shell) prompt() {
	fmt.Print(sh.PS1)
}

// processOutput will output user according to valid output, error
func processOutput(out string, err error) {
	if nil == err {
		fmt.Println(out)
	} else {
		fmt.Println(err)
	}
}

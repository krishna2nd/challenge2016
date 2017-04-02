// Package commands with 'set permissions' command implementation
package commands

import (
	"msg"
	"fmt"
	"bufio"
	"distributor"
)

// CmdSetPermission defined arguments and related methods
type CmdSetPermission struct {
	Command
	inputSubCmds []string
}

// NewCmdSetPermission new set permission command instance
func NewCmdSetPermission() *CmdSetPermission {
	var cmd = new(CmdSetPermission)
	cmd.Cmd = "permissions_for"
	return cmd
}

// Help to print help of set permission command
func (cp *CmdSetPermission) Help() {
	fmt.Println(`🔸  permissions_for -

  This allows to set permissions for distributor.
  eg:
  permissions_for DISTRIBUTOR1
	INCLUDE: INDIA
	INCLUDE: UNITEDSTATES
	EXCLUDE: KARNATAKA-INDIA
	EXCLUDE: CHENNAI-TAMILNADU-INDIA
	
    This allows DISTRIBUTOR1 to distribute in any city inside the United States and India,
	except cities in the state of Karnataka (in India) and the city of Chennai (in Tamil Nadu, India).
	
  permissions_for DISTRIBUTOR2 < DISTRIBUTOR1
	INCLUDE: INDIA
	EXCLUDE: TAMILNADU-INDIA
	
    DISTRIBUTOR2 can distribute the movie anywhere in INDIA, except inside TAMILNADU-INDIA
    and KARNATAKA-INDIA - DISTRIBUTOR2's permissions are always a subset of DISTRIBUTOR1's
    permissions. It's impossible/invalid for DISTRIBUTOR2 to have INCLUDE: CHINA,
    for example, because DISTRIBUTOR1 isn't authorized to do that in the first place.
  
  permissions_for DISTRIBUTOR3 < DISTRIBUTOR2 < DISTRIBUTOR1
	INCLUDE: HUBLI-KARNATAKA-INDIA

    DISTRIBUTOR2 cannot authorize DISTRIBUTOR3 with a region that they themselves do not have access to.
`)
}

// Parse to parse arguments
func (cp *CmdSetPermission) Parse(argString string, reader *bufio.Scanner) error {
	cp.Command.Parse(argString, reader)
	cp.inputSubCmds =  make([]string, 0)
	// Read input  sub commands
	for reader.Scan() {
		input := reader.Text()
		if msg.Empty == input {
			break
		}
		cp.inputSubCmds = append(cp.inputSubCmds, input)
	}
	return nil
}

// Verify to check the provided parameters are valid or not
func (cp *CmdSetPermission) Verify() error {
	if 0 == len(cp.Args) {
		return msg.ErrInvalidParams
	}
	return nil;
}

// Run to execute the command and provide result
func (cp *CmdSetPermission) Run() (string, error) {
	if (cp.isHelp()) {
		cp.Help();
		return msg.Empty, nil
	}
	switch len(cp.Args) {
	case 1: return distributor.SetIndividualPermission(cp.Args[0], cp.inputSubCmds);
	case 2: return distributor.SetLevelPermission(cp.Args[1], cp.Args[0], cp.inputSubCmds);
	case 3: return distributor.SetLevel2Permission(cp.Args[2], cp.Args[1], cp.Args[0], cp.inputSubCmds);
	}
	return msg.Empty, msg.ErrInvalidCommandExcec
}


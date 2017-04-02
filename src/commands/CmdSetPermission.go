// Package commands with 'set permissions' command implementation
package commands

import (
	"msg"
	"fmt"
)

// CmdSetPermission defined arguments and related methods
type CmdSetPermission struct {
	Command
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
	
  permissions_for DISTRIBUTOR2 < DISTRIBUTOR1
	INCLUDE: INDIA
	EXCLUDE: TAMILNADU-INDIA
  
  permissions_for DISTRIBUTOR3 < DISTRIBUTOR2 < DISTRIBUTOR1
	INCLUDE: HUBLI-KARNATAKA-INDIA

  This allows DISTRIBUTOR1 to distribute in any city inside the United States and India,
except cities in the state of Karnataka (in India) and the city of Chennai (in Tamil Nadu, India).
`)
}

// Parse to parse arguments
func (cp *CmdSetPermission) Parse(argString string) error {
	cp.Command.Parse(argString)
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
		return "", nil
	}
	switch len(cp.Args) {
	case 1: return cp.setIndividualPermission();
	case 2: return cp.setLevel1Permission();
	case 3: return cp.setLevel2Permission();
	}
	return "", msg.ErrInvalidCommandExcec
}

// setIndividualPermission to execute the command and provide result
func (cp *CmdSetPermission) setIndividualPermission() (string, error) {
	fmt.Println("Individual permission")
	return "", nil
}

// setLevel1Permission to execute the command and provide result
func (cp *CmdSetPermission) setLevel1Permission() (string, error) {
	fmt.Println("Level 1 permission")
	return "", nil
}

// setLevel2Permission to execute the command and provide result
func (cp *CmdSetPermission) setLevel2Permission() (string, error) {
	fmt.Println("Level2 permission")
	return "", nil
}
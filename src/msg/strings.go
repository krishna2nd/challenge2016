package msg

import (
	"errors"
	"fmt"
)

var (
	ErrCommandParsing           = errors.New("command: Error while parsing command and arguments")
	ErrInvalidParams            = errors.New("command: Invalid parameters please provide valid input")
	ErrInvalidCommand           = errors.New("command: Please input valid command")
	ErrInvalidCommandExcec      = errors.New("command: InputPlease input valid command")
	ErrNotFound                 = errors.New("Not found")
	
	DistrbuterNotFound          = "Distributer %v not found."
	ErrInheritPermission        = "Parent distributer %v does not have sufficient permissions for %v"
	Empty   = ""
	Space   = " "
	NewLine = "\n"
	Zero    = 0
	Comma   = ", "
	EndLine = '\n'
	Seperator = "<"
	CutSet = "\n\t\b\v"
	Success = "OK"
)

func ErrDistributerNotFound(msg string) error {
	return errors.New(fmt.Sprintf(DistrbuterNotFound, msg))
}
func ErrInheritPermissions(d1, d2 string) error {
	return errors.New(fmt.Sprintf(ErrInheritPermission, d1, d2))
}
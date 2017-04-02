package msg

import "errors"

var (
	ErrCommandParsing           = errors.New("command: Error while parsing command and arguments")
	ErrInvalidParams            = errors.New("command: Invalid parameters please provide valid input")
	ErrInvalidCommand           = errors.New("command: Please input valid command")
	ErrInvalidCommandExcec      = errors.New("command: InputPlease input valid command")
	ErrNotFound                 = errors.New("Not found")
	
	Empty   = ""
	Space   = " "
	NewLine = "\n"
	Zero    = 0
	Comma   = ", "
	EndLine = '\n'
	Seperator = "<"
)
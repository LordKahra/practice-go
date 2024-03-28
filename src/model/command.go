package model

import "errors"

type Command struct {
	Type       string         `json:"type"`
	Credential HackCredential `json:"credential"`
	Arguments  []string       `json:"arguments"`
}

func (command Command) Handle() (string, error) {

	switch command.Type {
	case "help":
		return command.handleHelp()
	default:
		return "Unsupported type.", errors.New("unsupported type")
	}
}

func (command Command) handleHelp() (string, error) {
	// Break open the arguments.

	return "not yet implemented.", nil
}

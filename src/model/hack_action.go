package model

import (
	"database/sql"
	"errors"
)

type HackAction struct {
	Type        string `json:"type"`
	CharacterID int64  `json:"character_id"`
	ServerID    int64  `json:"server_id"`
}

func (action HackAction) Handle(db *sql.DB) (string, error) {

	switch action.Type {
	case "gain":
		return action.handleGainAction()
	case "print":
		return "not yet implemented", nil
	case "upload":
		return action.handleUploadAction()
	case "download":
		return action.handleDownloadAction()
	default:
		return "Unsupported type.", errors.New("unsupported type")
	}
}

func (action HackAction) handleGainAction() (string, error) {
	return "not yet implemented", nil
}

func (action HackAction) handleUploadAction() (string, error) {
	return "not yet implemented", nil
}

func (action HackAction) handleDownloadAction() (string, error) {
	return "not yet implemented", nil
}

package model

type HackCharacterFile struct {
	ID          int64  `json:"id"`
	CharacterID string `json:"character_id"`
	Filename    string `json:"filename"`
	Extension   string `json:"extension"`
	Data        string `json:"data"`
}

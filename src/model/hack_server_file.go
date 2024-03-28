package model

type HackServerFile struct {
	ID        int64  `json:"id"`
	ServerID  string `json:"server_id"`
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
	Data      string `json:"data"`
}

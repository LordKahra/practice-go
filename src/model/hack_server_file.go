package model

type HackServerFile struct {
	ID        int64  `json:"id"`
	ServerID  int64  `json:"server_id"`
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
	Data      string `json:"data"`
	IntelID   int64  `json:"intel_id"`
}

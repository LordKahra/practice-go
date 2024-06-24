package model

type Site struct {
	ID        int64  `json:"id"`
	NetworkID int64  `json:"network_id"`
	Address   string `json:"address"`
	Content   string `json:"content"`
}

package model

type HackCredential struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	ServerID int64  `json:"server_id"`
	IsActive bool   `json:"is_active"`
}

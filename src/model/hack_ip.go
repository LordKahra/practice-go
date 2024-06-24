package model

type HackIP struct {
	IPv4          string `json:"ipv4"`
	ServerId      int64  `json:"server_id"`
	ServerName    string `json:"server_name"`
	EffectiveDate int64  `json:"effective_date"`
	Status        string `json:"status"`
}

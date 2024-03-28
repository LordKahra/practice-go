package model

type HackServer struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	IPv4    string `json:"ipv4"`
	Address string `json:"address"`
}

package model

type HackServer struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	IPv4            string `json:"ipv4"`
	Address         string `json:"address"`
	CharacterID     int64  `json:"character_id"`
	Tags            string `json:"tags"`
	IPEffectiveDate int64  `json:"ip_effective_date"`
}

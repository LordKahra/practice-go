package model

type HackIntel struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	TypeId    int64  `json:"type_id"`
	TypeName  string `json:"type_name"`
	Target    string `json:"target"`
	IsVisible bool   `json:"is_visible"`
}

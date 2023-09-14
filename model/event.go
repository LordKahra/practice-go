package model

type Event struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	ChapterId    int64  `json:"chapter_id"`
	LocationId   int64  `json:"location_id"`
	DateStart    int64  `json:"date_start"`
	DateEnd      int64  `json:"date_end"`
	LinkFacebook string `json:"link_facebook"`
	LinkGoogle   string `json:"link_google"`
	LinkSite     string `json:"link_site"`
}

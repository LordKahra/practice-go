package model

import (
	"database/sql"
	"time"
)

type Chapter struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	DisplayName    string         `json:"display_name"`
	SystemID       int64          `json:"system_id"`
	OrganizationID int64          `json:"organization_id"`
	AddressID      int64          `json:"address_id"`
	LinkFacebook   sql.NullString `json:"link_facebook"`
	LinkSite       sql.NullString `json:"link_site"`
	LinkDiscord    sql.NullString `json:"link_discord"`
	Description    sql.NullString `json:"description"`
	Validated      time.Time      `json:"validated"`
}

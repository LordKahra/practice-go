package model

import (
	"database/sql"
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
)

type Address struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Address1    string          `json:"address_1"`
	Address2    sql.NullString  `json:"address_2"`
	City        string          `json:"city"`
	State       string          `json:"state"`
	Zip         string          `json:"zip"`
	Description sql.NullString  `json:"description"`
	LinkMaps    sql.NullString  `json:"link_maps"`
	Latitude    sql.NullFloat64 `json:"latitude"`
	Longitude   sql.NullFloat64 `json:"longitude"`
}

func GetAddressString(address Address) string {
	var str = address.Address1 + ", "
	if address.Address2.Valid {
		str += address.Address2.String + ", "
	}
	str += address.City + ", " + address.State + " " + address.Zip
	return str
}

func GetAddressGeocode(address Address) (*geo.Location, error) {
	geocoder := openstreetmap.Geocoder()
	location, err := geocoder.Geocode(GetAddressString(address))

	// Pass straight through, we're just making it easy for Address structs.
	return location, err
}

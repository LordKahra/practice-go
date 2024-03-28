package reflection

import (
	"database/sql"
	"encoding/json"
)

// Overrides.

type NullString struct {
	sql.NullString
}

func (x *NullString) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		x.Valid = true
		x.String = ""
		//return []byte("null"), nil
	}
	return json.Marshal(x.String)
}

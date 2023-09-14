package database

import (
	. "practice-go/model"
)

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// GET

func GetEvents(db *sql.DB, where string) ([]Event, error) {
	query := "SELECT id, name, chapter_id, location_id, " +
		"date_start, date_end, link_fb, link_google, link_site " +
		"FROM events"
	if where != "" {
		query += " WHERE " + where
	}

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		// Prep potentially null objects.
		var LocationID sql.NullInt64
		var LinkFacebook sql.NullString
		var LinkGoogle sql.NullString
		var LinkSite sql.NullString

		// Attempt to scan the row into the event.
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.ChapterId, &LocationID,
			&event.DateStart, &event.DateEnd,
			&LinkFacebook, &LinkGoogle, &LinkSite)
		if err != nil {
			return nil, err
		}

		// Process nullables.
		if LocationID.Valid {
			event.LocationId = LocationID.Int64
		}
		if LinkFacebook.Valid {
			event.LinkFacebook = LinkFacebook.String
		}
		if LinkGoogle.Valid {
			event.LinkGoogle = LinkGoogle.String
		}
		if LinkSite.Valid {
			event.LinkSite = LinkSite.String
		}

		events = append(events, event)
	}

	// Done.
	return events, rows.Err()
}

func GetChapters(db *sql.DB) ([]Chapter, error) {
	rows, err := db.Query(
		"SELECT id, name, display_name, system_id, organization_id, address_id, " +
			"link_fb, link_site, link_discord, description, validated " +
			"FROM chapters")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chapters []Chapter
	for rows.Next() {

		// Prep potentially null objects.

		// Attempt to scan the row into the event.
		var chapter Chapter
		if err := rows.Scan(&chapter.ID, &chapter.Name, &chapter.DisplayName,
			&chapter.SystemID, &chapter.OrganizationID, &chapter.AddressID,
			&chapter.LinkFacebook, &chapter.LinkSite, &chapter.LinkDiscord,
			&chapter.Description, &chapter.Validated,
		); err != nil {
			return nil, err
		}
		chapters = append(chapters, chapter)
	}

	// Done.
	return chapters, rows.Err()
}

// POST - CREATION

func CreateEvent(db *sql.DB, event Event) (int64, error) {
	// Create the query.
	query := `INSERT INTO events (name, chapter_id, location_id, 
                    date_start, date_end,
                    link_fb, link_google, link_site
                    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	// Run the query.
	result, err := db.Exec(query, event.Name,
		event.ChapterId, event.LocationId, event.DateStart, event.DateEnd,
		event.LinkFacebook, event.LinkGoogle, event.LinkSite)

	if err != nil {
		return 0, err
	}

	eventID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	// Done!
	return eventID, nil
}

// PUT - UPDATES

func UpdateEvent(db *sql.DB, params map[string]string) {

}

// DELETE

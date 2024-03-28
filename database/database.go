package database

import (
	"errors"
	. "practice-go/model"
)

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

////////////////////////////////
//// GET ///////////////////////
////////////////////////////////

//// HACKING ///////////////////

func GetHackCharacterFile(db *sql.DB, fileId int64) (HackCharacterFile, error) {
	var file HackCharacterFile

	query :=
		`SELECT id, character_id, filename, extension, data FROM hack_character_files
			WHERE id = ?`

	rows, err := db.Query(query, fileId)

	if err != nil {
		return file, err
	}
	defer rows.Close()

	for rows.Next() {
		// Prep potentially null objects.
		// None, continue.

		// Scan the row into the HackServer.

		err = rows.Scan(&file.ID, &file.CharacterID, &file.Filename, &file.Extension, &file.Data)
		if err != nil {
			return file, err
		}

		// Done. Return.
		return file, nil
	}
	// File not found.
	return file, errors.New("file not found")
}

func GetHackCharacterCredentials(db *sql.DB, characterId int64) ([]HackCredential, error) {
	query :=
		`SELECT credential.id, credential.username, credential.password, credential.server_id 
			FROM hack_credentials credential 
			LEFT JOIN hav.hack_r_character_credentials char_creds on credential.id = char_creds.credential_id 
			WHERE char_creds.character_id = ?`

	rows, err := db.Query(query, characterId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var credentials []HackCredential
	for rows.Next() {
		// Prep potentially null objects.
		// None, continue.

		// Scan the row into the HackServer.
		var credential HackCredential
		err = rows.Scan(&credential.ID, &credential.Username, &credential.Password, &credential.ServerID)
		if err != nil {
			return nil, err
		}

		// Process nullables.
		// None to process.

		credentials = append(credentials, credential)
	}

	// Done.
	return credentials, rows.Err()
}

func GetHackCharacterServers(db *sql.DB, characterId int64) ([]HackServer, error) {
	query :=
		`SELECT server.id, server.name, server.ipv4, server.address 
			FROM hack_servers server 
			LEFT JOIN hav.hack_r_character_servers char_servers on server.id = char_servers.server_id 
			WHERE char_servers.character_id = ?`

	rows, err := db.Query(query, characterId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var servers []HackServer
	for rows.Next() {
		// Prep potentially null objects.
		// None, continue.

		// Scan the row into the HackServer.
		var server HackServer
		err = rows.Scan(&server.ID, &server.Name, &server.IPv4, &server.Address)
		if err != nil {
			return nil, err
		}

		// Process nullables.
		// None to process.

		servers = append(servers, server)
	}

	// Done.
	return servers, rows.Err()
}

func GetHackServers(db *sql.DB, where string) ([]HackServer, error) {
	query := "SELECT id, name, ipv4, address FROM hack_servers"
	if where != "" {
		query += " WHERE " + where
	}

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var servers []HackServer
	for rows.Next() {
		// Prep potentially null objects.
		// None, continue.

		// Scan the row into the HackServer.
		var server HackServer
		err = rows.Scan(&server.ID, &server.Name, &server.IPv4, &server.Address)
		if err != nil {
			return nil, err
		}

		// Process nullables.
		// None to process.

		servers = append(servers, server)
	}

	// Done.
	return servers, rows.Err()
}

func GetHackServerFile(db *sql.DB, fileId int64) (HackServerFile, error) {
	var file HackServerFile

	query :=
		`SELECT id, server_id, filename, extension, data FROM hack_server_files
			WHERE id = ?`

	rows, err := db.Query(query, fileId)

	if err != nil {
		return file, err
	}
	defer rows.Close()

	for rows.Next() {
		// Prep potentially null objects.
		// None, continue.

		// Scan the row into the HackServer.

		err = rows.Scan(&file.ID, &file.ServerID, &file.Filename, &file.Extension, &file.Data)
		if err != nil {
			return file, err
		}

		// Done. Return.
		return file, nil
	}
	// File not found.
	return file, errors.New("file not found")
}

//// NON-HACKING ///////////////////

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

////////////////////////////////
//// POST - CREATION ///////////
////////////////////////////////

// HACKING ////

func CreateHackUpload(db *sql.DB, serverId int64, file HackCharacterFile) (int64, error) {
	// Create the query.
	query := `INSERT INTO hack_server_files (server_id, filename, extension, data) VALUES (?, ?, ?, ?)`

	// Run the query.
	result, err := db.Exec(query, serverId, file.Filename, file.Extension, file.Data)

	if err != nil {
		return 0, err
	}

	fileID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	// Done!
	return fileID, nil
}

func CreateHackDownload(db *sql.DB, characterId int64, file HackServerFile) (int64, error) {
	// Create the query.
	query := `INSERT INTO hack_character_files (character_id, filename, extension, data) VALUES (?, ?, ?, ?)`

	// Run the query.
	result, err := db.Exec(query, characterId, file.Filename, file.Extension, file.Data)

	if err != nil {
		return 0, err
	}

	fileID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	// Done!
	return fileID, nil
}

// HACKING - LOGGING IN

func HackConnectToServer(db *sql.DB, credential HackCredential) (HackServer, error) {
	var server HackServer

	query := `SELECT server.id, server.name, server.ipv4, server.address
				FROM hack_servers server
				LEFT JOIN hav.hack_credentials creds on server.id = creds.server_id
				WHERE creds.username = ? AND creds.password = ? AND server.id = ?`

	rows, err := db.Query(query, credential.Username, credential.Password, credential.ServerID)

	if err != nil {
		return server, err
	}
	defer rows.Close()

	for rows.Next() {
		// Prep potentially null objects.
		// None, continue.

		// Scan the row into the HackServer.

		err = rows.Scan(&server.ID, &server.Name, &server.IPv4, &server.Address)
		if err != nil {
			return server, err
		}

		// Done. Return.
		return server, nil
	}
	// File not found.
	return server, errors.New("server not found")
}

// NON-HACKING

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

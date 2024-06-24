package database

import (
	"errors"
	. "practice-go/model"
)

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

////
//// CONSTANTS?
////

const (
	NOT_FOUND string = "no records found"
)

////////////////////////////////
//// GET ///////////////////////
////////////////////////////////

//// HACKING ///////////////////

func GetHackCharacterIntel(db *sql.DB, characterId int64) ([]HackIntel, []HackServer, []HackCredential, []HackIP, []int64, []int64, []int64, []int64, []int64, []string, error) {
	query := `SELECT intel.id as id, intel.name as name, intel.type_id as type_id, intel.type_name as type_name, intel.target as target, intel.is_visible as is_visible
				FROM hack_character_intel_details intel
				WHERE intel.character_id = ?
				ORDER BY intel.type_id, intel.target`

	rows, err := db.Query(query, characterId)

	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	defer rows.Close()

	var intels []HackIntel
	hasServers := false
	hasCredentials := false
	hasIPs := false
	var servers []HackServer
	var credentials []HackCredential
	var ips []HackIP
	var dictionaryServerIds []int64
	var rainbowServerIds []int64
	var sprays []int64
	var malware []int64
	var research []int64
	var emails []string

	for rows.Next() {
		var intel HackIntel
		if err := rows.Scan(
			&intel.ID, &intel.Name, &intel.TypeId, &intel.TypeName, &intel.Target, &intel.IsVisible,
		); err != nil {
			return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
		}
		intels = append(intels, intel)

		// Based on type, add to the appropriate bin.
		switch typeName := intel.TypeName; typeName {
		case "server":
			hasServers = true
		case "credential":
			hasCredentials = true
		case "ip_address":
			hasIPs = true
		case "dictionary":
		case "rainbow":
		case "spray":
		case "malware":
		case "research":
		case "email_address":
		}
	}

	// Retrieval information gathered. Fetch individual values.
	if hasServers {
		servers, err = GetHackCharacterServers(db, characterId)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
		}
	}
	if hasCredentials {
		credentials, err = GetHackCharacterCredentials(db, characterId)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
		}
	}
	if hasIPs {
		ips, err = GetHackCharacterIPv4s(db, characterId)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
		}
	}

	// Done.
	return intels, servers, credentials, ips, dictionaryServerIds, rainbowServerIds, sprays, malware, research, emails, rows.Err()
}

func GetHackCharacterCredentials(db *sql.DB, characterId int64) ([]HackCredential, error) {
	query := `SELECT cred.id as id, cred.username as username, cred.password as password, cred.server_id as server_id, cred.is_active as is_active
				FROM hack_character_intel_details intel
				LEFT JOIN hack_credentials cred ON cred.id = intel.target
				WHERE intel.character_id = ? AND intel.type_name = 'credential'`

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
		err = rows.Scan(&credential.ID, &credential.Username, &credential.Password, &credential.ServerID, &credential.IsActive)
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

func GetHackCharacterIPv4s(db *sql.DB, characterId int64) ([]HackIP, error) {
	query := `SELECT ip.ipv4 as ipv4, ip.server_id as server_id, ip.server_name as server_name, ip.effective_date as effective_date, ip.status as status
				FROM hack_ip_details ip
				INNER JOIN hack_character_intel_details intel ON intel.target = ip.ipv4
				WHERE intel.character_id = ? AND intel.type_name = 'ip_address'
				ORDER BY ip.status DESC, ip.ipv4 asc`

	rows, err := db.Query(query, characterId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ips []HackIP
	for rows.Next() {
		var ip, err = scanCurrentIPRow(rows)
		if err != nil {
			return nil, err
		}

		ips = append(ips, ip)
	}

	// Done.
	return ips, rows.Err()
}

func GetHackCharacterServers(db *sql.DB, characterId int64) ([]HackServer, error) {
	query := `SELECT server.id as id, server.name as name, ip.ipv4 as ipv4, server.address as address,
       			server.character_id as character_id, server.tags as tags,
       			ip.effective_date as ip_effective_date
				FROM hack_character_intel_details intel
				LEFT JOIN hack_ip_details ip ON intel.target = ip.ipv4
				LEFT JOIN hack_server_details server ON ip.server_id = server.id
				WHERE intel.character_id = ? AND intel.type_name = 'ip_address' AND ip.status = 'online'
				ORDER BY ip.status DESC, server.name asc`

	rows, err := db.Query(query, characterId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var servers []HackServer
	for rows.Next() {
		var server, err = scanCurrentServerRow(rows)
		if err != nil {
			return nil, err
		}

		servers = append(servers, server)
	}

	// Done.
	return servers, rows.Err()
}

func GetHackServers(db *sql.DB, where string) ([]HackServer, error) {
	query := "SELECT id, name, ipv4, address, character_id, tags, ip_effective_date FROM hack_server_details"
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
		var server, err = scanCurrentServerRow(rows)
		if err != nil {
			return nil, err
		}

		servers = append(servers, server)
	}

	// Done.
	return servers, rows.Err()
}

func GetHackServerFile(db *sql.DB, fileId int64) (HackServerFile, error) {
	var file HackServerFile

	query :=
		`SELECT id, server_id, filename, extension, data, intel_id FROM hack_server_files
			WHERE id = ?`

	rows, err := db.Query(query, fileId)

	if err != nil {
		return file, err
	}
	defer rows.Close()

	for rows.Next() {
		// Prep potentially null objects.
		var intelId sql.NullInt64

		// Scan the row into the HackServer.

		err = rows.Scan(&file.ID, &file.ServerID, &file.Filename, &file.Extension, &file.Data, &intelId)
		if err != nil {
			return file, err
		}

		// Process nullables.
		if intelId.Valid {
			file.IntelID = intelId.Int64
		}

		// Done. Return.
		return file, nil
	}
	// File not found.
	return file, errors.New("file not found")
}

//// NON-HACKING ///////////////////

func GetCharacter(db *sql.DB, characterId int64) (Character, error) {
	var character Character

	// Query for a single row.
	row := db.QueryRow("SELECT id, name, player_id FROM characters where id = ?", characterId)

	// Scan into your struct.
	err := row.Scan(&character.ID, &character.Name, &character.PlayerID)
	if err != nil {
		return character, err
	}

	return character, nil
}

func GetSite(db *sql.DB, siteId int64) (Site, error) {
	var site Site

	// Query for a single row.
	row := db.QueryRow("SELECT id, network_id, address, content FROM sites WHERE id = ?", siteId)

	// Scan into your struct.
	err := row.Scan(&site.ID, &site.NetworkID, &site.Address, &site.Content)
	if err != nil {
		return site, err
	}

	return site, nil
}

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

func GetSites(db *sql.DB, where string) ([]Site, error) {
	query := "SELECT id, network_id, address, content FROM sites"
	if where != "" {
		query += " WHERE " + where
	}

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sites []Site
	for rows.Next() {
		var site, err = scanCurrentSiteRow(rows)
		if err != nil {
			return nil, err
		}

		sites = append(sites, site)
	}

	// Done.
	return sites, rows.Err()
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

func HackTransferFile(db *sql.DB, targetServerId int64, file HackServerFile) (int64, error) {
	// Create variables.
	var result sql.Result
	var err error

	// Create and run the query.
	if file.IntelID == 0 {
		query := `INSERT INTO hack_server_files (server_id, filename, extension, data) VALUES (?, ?, ?, ?)`

		result, err = db.Exec(query, targetServerId, file.Filename, file.Extension, file.Data)
	} else {
		query := `INSERT INTO hack_server_files (server_id, filename, extension, data, intel_id) VALUES (?, ?, ?, ?, ?)`

		result, err = db.Exec(query, targetServerId, file.Filename, file.Extension, file.Data, file.IntelID)
	}
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

	query := `SELECT server.id, server.name, server.ipv4, server.address,
				server.character_id, server.tags, server.ip_effective_date
				FROM hack_server_details server
				LEFT JOIN hav.hack_credentials creds on server.id = creds.server_id
				WHERE creds.username = ? AND creds.password = ? AND server.id = ?`

	rows, err := db.Query(query, credential.Username, credential.Password, credential.ServerID)

	if err != nil {
		return server, err
	}
	defer rows.Close()

	for rows.Next() {
		var server, err = scanCurrentServerRow(rows)
		if err != nil {
			return server, err
		}

		// Done. Return.
		return server, nil
	}
	// File not found.
	return server, errors.New("server not found")
}

func GetHackFilesByCredential(db *sql.DB, serverId int64, credentialId int64, ipv4 string) ([]HackServerFile, error) {
	query := `SELECT 
					file.id as id, file.server_id as server_id, file.filename as filename, 
					file.extension as extension, file.data as data, file.intel_id as intel_id
				FROM hack_server_files file
				INNER JOIN hack_credentials creds on file.server_id = creds.server_id
				INNER JOIN hack_server_details server on server.id = file.server_id
				WHERE creds.id = ? AND server.ipv4 = ? AND server.id = ?
				ORDER BY file.filename`

	rows, err := db.Query(query, credentialId, ipv4, serverId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hackServerFiles []HackServerFile
	for rows.Next() {
		var file, err = scanCurrentFileRow(rows)
		if err != nil {
			return nil, err
		}
		hackServerFiles = append(hackServerFiles, file)
	}

	// Done.
	return hackServerFiles, rows.Err()
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

////
// SCANNING

func scanCurrentFileRow(rows *sql.Rows) (HackServerFile, error) {
	// Prep potentially null objects.
	var IntelID sql.NullInt64

	// Scan the row into the file.
	var file HackServerFile
	var err error
	err = rows.Scan(&file.ID, &file.ServerID, &file.Filename,
		&file.Extension, &file.Data, &IntelID)
	if err != nil {
		return file, err
	}

	// Process nullables.
	if IntelID.Valid {
		file.IntelID = IntelID.Int64
	}

	return file, nil
}

func scanCurrentIPRow(rows *sql.Rows) (HackIP, error) {
	// Scan the row into a HackIP.
	var IPv4 HackIP
	var err error
	err = rows.Scan(
		&IPv4.IPv4, &IPv4.ServerId, &IPv4.ServerName,
		&IPv4.EffectiveDate, &IPv4.Status)
	if err != nil {
		return IPv4, err
	}

	// Done.
	return IPv4, nil
}

func scanCurrentServerRow(rows *sql.Rows) (HackServer, error) {
	// Prep potentially null objects.
	var Address sql.NullString
	var CharacterID sql.NullInt64
	var Tags sql.NullString

	// Scan the row into the HackServer.
	var server HackServer
	var err error
	err = rows.Scan(
		&server.ID, &server.Name, &server.IPv4, &Address,
		&CharacterID, &Tags, &server.IPEffectiveDate,
	)
	if err != nil {
		return server, err
	}

	// Process nullables.
	if Address.Valid {
		server.Address = Address.String
	}
	if CharacterID.Valid {
		server.CharacterID = CharacterID.Int64
	}
	if Tags.Valid {
		server.Tags = Tags.String
	}

	return server, nil
}

func scanCurrentSiteRow(rows *sql.Rows) (Site, error) {
	// No null objects.
	// Scan the row into the Site.
	var site Site
	var err error
	err = rows.Scan(&site.ID, &site.NetworkID, &site.Address, &site.Content)
	if err != nil {
		return site, err
	}

	// Done. No nullables to process.
	return site, nil
}

package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connectToDatabase() (*sql.DB, error) {
	user := "root"
	password := "insecure"
	host := "localhost"
	port := "8080"
	database := "southern_larp"

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

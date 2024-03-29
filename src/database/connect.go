package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() (*sql.DB, error) {
	user := "root"
	//password := "insecure"
	password := "insecure"
	host := "localhost"
	port := "3306"
	database := "hav"

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&collation=utf8mb4_0900_ai_ci", user, password, host, port, database)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

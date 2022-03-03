package database

import (
	"database/sql"
	"log"
)

func Create(db *sql.DB) (err error) {
	userTable := "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, app TEXT, username TEXT);"
	statement, err := db.Prepare(userTable)

	if err != nil {
		log.Println(err.Error())

	}

	if _, err = statement.Exec(); err != nil {
		log.Println(err.Error())

	}

	return statement.Close()
}

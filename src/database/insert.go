package database

import (
	"database/sql"
	"log"
)

func Insert(db *sql.DB, id string, client string, username string) (err error) {
	insertUser := "INSERT INTO stats (client, username) VALUES ($1, $2) RETURNING (client, username);"
	rows, err := db.Query(insertUser, client, username)

	if err == nil {
		log.Println("The data has been updated, now it is:", id, client, username)

		return rows.Close()
	}

	log.Println(err.Error())
	return rows.Close()
}

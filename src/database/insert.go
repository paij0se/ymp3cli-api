package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Insert(db *sql.DB, id string, app string, username string) (err error) {
	insertUser := "INSERT INTO users (app, username) VALUES ($1 $2) RETURNING (app, username);"
	rows, err := db.Query(insertUser, app, username)

	if err == nil {
		log.Println("The data has been updated, now it is:", id, app, username)

	}

	log.Println(err.Error())
	return rows.Close()
}

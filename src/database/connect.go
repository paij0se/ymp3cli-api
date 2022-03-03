package database

import (
	"database/sql"
	"log"
	"os"
)

func Connect() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", os.Getenv("DB_URI"))

	if err != nil {
		log.Println(err.Error())

	}

	if err = Create(db); err != nil {
		log.Println(err.Error())

	}

	return
}

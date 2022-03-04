package database

import (
	"database/sql"
	"log"

	"github.com/paij0se/ymp3cli-api/src/interfaces"
)

func Query(db *sql.DB, API *interfaces.Ymp3cli) (err error) {
	row, err := db.Query("SELECT * FROM users ORDER BY id DESC LIMIT 1;")

	if err != nil {
		log.Println(err.Error())

		return
	}

	if row.Next() {
		if err = row.Scan(&API.Id, &API.Client, &API.Username); err != nil {
			log.Println(err.Error())

		}

	}

	return row.Close()
}

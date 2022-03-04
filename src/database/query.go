package database

import (
	"database/sql"
	"log"

	"github.com/paij0se/ymp3cli-api/src/interfaces"
)

func Query(db *sql.DB, API *[]interfaces.Ymp3cli) (err error) {
	row, err := db.Query("SELECT * FROM stats ORDER BY id DESC LIMIT 1;")

	if err != nil {
		log.Println(err.Error())

		return
	}

	for row.Next() {
		var ymp3cli interfaces.Ymp3cli

		if err = row.Scan(&ymp3cli.Id, &ymp3cli.Client, &ymp3cli.Username); err != nil {
			log.Println(err.Error())
		}

		*API = append(*API, ymp3cli)
	}

	log.Println(API)

	return row.Close()
}

package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/paij0se/ymp3cli-api/src/interfaces"
)

func Query(db *sql.DB, id *int, ymp3cli *interfaces.Ymp3cli) (err error) {
	row, err := db.Query("SELECT * FROM users ORDER BY id DESC LIMIT 1;")

	if err != nil {
		log.Println(err.Error())

	} else if row.Next() {
		if err = row.Scan(id, ymp3cli.App, ymp3cli.Username); err != nil {
			log.Println(err.Error())

		}
	}

	return row.Close()
}

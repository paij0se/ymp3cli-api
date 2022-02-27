package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Db(id string, user string) {
	sqliteDatabase, err := sql.Open("sqlite3", "./src/database/database.db")
	checkErr(err)
	createTable(sqliteDatabase)
	insertData(sqliteDatabase, id, user)
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `
	CREATE TABLE IF NOT EXISTS users ( 
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		User TEXT NOT NULL
	);
	`
	statement, err := db.Prepare(createStudentTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertData(db *sql.DB, id string, user string) {
	log.Println("Inserting Data")
	insertStudentSQL := "INSERT INTO users (User) VALUES (?);"
	statement, err := db.Prepare(insertStudentSQL)
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(user) // Execute SQL Query
	if err != nil {
		log.Fatalln(err.Error())
	}

}
func DisplayLastUser(db *sql.DB) {
	row, err := db.Query("SELECT * FROM users ORDER BY id DESC LIMIT 1")
	checkErr(err)
	defer row.Close()
	for row.Next() {
		var id int
		var user string
		err = row.Scan(&id, &user)
		checkErr(err)
		log.Println(id, user)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

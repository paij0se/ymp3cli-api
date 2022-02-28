package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

type user struct {
	User string
}

func createDB(db *sql.DB) {
	// create users table if not exists
	createTableUsers := "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR NOT NULL);"
	statement, err := db.Prepare(createTableUsers)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	log.Println("Created table users")
}

func insertData(db *sql.DB, id string, user string) {
	log.Println("Inserting data")
	insertUser := `
	INSERT INTO users (name) VALUES ($1);
	`
	statement, err := db.Prepare(insertUser)
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(user) // Execute SQL Query
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func Db(id string, user string) {
	postgres, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	createDB(postgres)
	insertData(postgres, id, user)
}

func postDataUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user user
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(reqBody, &user)
	u := user.User
	Db("0001", u)
	c.JSON(200, gin.H{
		"message": u,
	})

}

func displayUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	postgres, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	row, err := postgres.Query("SELECT * FROM users ORDER BY id DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var name string
		err = row.Scan(&id, &name)
		if err != nil {
			log.Printf("Error scanning row: %q", err)
		}
		c.JSON(200, gin.H{
			"id":       id,
			"lastUser": name,
		})
	}

}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.New()
	router.Use(CORSMiddleware())
	router.POST("/user", postDataUser)
	router.GET("/", displayUser)
	router.Use(gin.Logger())
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT=8080 go run main.go")
	}
	router.Run(":" + port)
}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

type name struct {
	Name     string
	Password string
}

func createDB(db *sql.DB) {
	// create users table if not exists
	createTableUsers := "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY,name TEXT);"
	statement, err := db.Prepare(createTableUsers)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	log.Println("Created table users")
}

func insertData(db *sql.DB, id string, name string) {
	log.Println("Inserting data")
	insertUser := "INSERT INTO users (name) VALUES ($1) RETURNING name;"
	err := db.QueryRow(insertUser, name).Scan(&id)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("New record ID is:", id, name)
}

func Db(id string, name string) {
	postgres, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
	}
	createDB(postgres)
	// insert data
	insertData(postgres, id, name)

}

func postDataUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user name
	var password name
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(reqBody, &user)
	json.Unmarshal(reqBody, &password)
	if password.Name != "d24250200f82994615e9f4c7336c4aa4457e28d39491cce206f60f89b73594ea98a047edfd8e68afe7705a7b2098cd81cc14043cb52e87c391b8e68c9f65c354" {
		c.JSON(401, gin.H{
			"message": "🚫",
		})
	}
	Db("0001", user.Name)
	c.JSON(200, gin.H{
		"message": "added",
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
		log.Println(err)
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
		log.Println("$PORT must be set")
	}
	router.Run(":" + port)
}

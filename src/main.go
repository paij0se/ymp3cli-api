package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

type name struct {
	Name string
}

// TODO: Organize this shit

var (
	lastReq = uint64(time.Now().UnixMilli())
	cPeriod = uint64(60 * 60 * 1000)

	reqList = make(map[string]uint64)
)

func createDB(db *sql.DB) {
	// create users table if not exists
	createTableUsers := "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY,name TEXT);"
	statement, err := db.Prepare(createTableUsers)
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	log.Println("Created table users")
}

func insertData(db *sql.DB, id string, name string) {
	log.Println("Inserting data")
	insertUser := "INSERT INTO users (name) VALUES ($1) RETURNING name;"
	in, err := db.Query(insertUser, name)
	defer in.Close()
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
	// close conection
	defer postgres.Close()

}

func postDataUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user name
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	if len(user.Name) == 0 {
		user.Name = "Anonymous"
	}
	json.Unmarshal(reqBody, &user)
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
	row, err := postgres.Query("SELECT * FROM users ORDER BY id DESC LIMIT 1;")
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

func RateLimit(delay uint64) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		reqUrl := ctx.Request.URL.String() + ctx.ClientIP()
		dateNow := uint64(time.Now().UnixMilli())

		// clear memory.
		if (dateNow - lastReq) > cPeriod {
			reqList = make(map[string]uint64)

		}

		lastReq = dateNow

		if value, key := reqList[reqUrl]; key && dateNow < value {
			ctx.AbortWithStatusJSON(429, gin.H{
				"message": "429 - Too Many Requests.",
			})

			return
		}

		reqList[reqUrl] = (dateNow + delay)
		ctx.Next()
	}

}

func main() {
	router := gin.New()
	router.Use(CORSMiddleware())
	router.POST("/user", RateLimit(20000), postDataUser)
	router.GET("/", RateLimit(2000), displayUser)
	router.Use(gin.Logger())
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT must be set")
	}
	router.Run(":" + port)
}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/paij0se/ymp3cli-api/src/database"
)

type user struct {
	User string
}

func postUser(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusCreated)
	var user user
	reqBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Fprintf(c.Response(), "Error")
	}
	json.Unmarshal([]byte(reqBody), &user)
	u := user.User
	database.Db("0001", u)
	sqliteDatabase, err := sql.Open("sqlite3", "./src/database/database.db")
	if err != nil {
		fmt.Fprintf(c.Response(), "Error")
	}
	row, err := sqliteDatabase.Query("SELECT * FROM users ORDER BY id DESC LIMIT 1")
	if err != nil {
		fmt.Fprintf(c.Response(), "Error")
	}
	defer row.Close()
	for row.Next() {
		var id int
		var user string
		err = row.Scan(&id, &user)
		if err != nil {
			fmt.Fprintf(c.Response(), "Error")
		}
		fmt.Fprintf(c.Response(), `{"id": %d, "last-user": "%s"}`, id, user)
	}
	return nil

}

func displayUser(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusCreated)
	sqliteDatabase, err := sql.Open("sqlite3", "./src/database/database.db")
	if err != nil {
		fmt.Fprintf(c.Response(), "Error")
	}
	row, err := sqliteDatabase.Query("SELECT * FROM users ORDER BY id DESC LIMIT 1")
	if err != nil {
		fmt.Fprintf(c.Response(), "Error")
	}
	defer row.Close()
	for row.Next() {
		var id int
		var user string
		err = row.Scan(&id, &user)
		if err != nil {
			fmt.Fprintf(c.Response(), "Error")
		}
		fmt.Fprintf(c.Response(), `{"id": %d, "last-user": "%s"}`, id, user)
	}
	return nil
}

func main() {
	e := echo.New()
	e.GET("/", displayUser)
	e.POST("/user", postUser)
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "5000"
	}
	fmt.Printf("server on port: %s", port)
	e.Logger.Fatal(e.Start(":" + port))

}

package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/paij0se/ymp3cli-api/src/database"
	"github.com/paij0se/ymp3cli-api/src/interfaces"
)

func UpdateData(ctx *gin.Context) {
	var ymp3cli interfaces.Ymp3cli

	reqBody, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		log.Println(err.Error())

		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "500 - Internal Server Error.",
			"error":   err.Error(),
		})

		return
	}

	if err = json.Unmarshal(reqBody, &ymp3cli); err != nil {
		log.Println(err.Error())

		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "500 - Internal Server Error.",
			"error":   err.Error(),
		})

		return
	}

	if ymp3cli.Client == "" || ymp3cli.Username == "" {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "400 - Bad Request.",
		})

		return
	}

	db, err := database.Connect()

	if err != nil {
		log.Println(err.Error())

		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "500 - Internal Server Error.",
			"error":   err.Error(),
		})

		return
	}

	if err = database.Insert(db, "0001", ymp3cli.Client, ymp3cli.Username); err != nil {
		log.Println(err.Error())

		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "500 - Internal Server Error.",
			"error":   err.Error(),
		})

		return
	}

	if err = db.Close(); err != nil {
		log.Println(err.Error())

	}

	ctx.JSON(200, gin.H{
		"message": "200 - OK",
	})
}

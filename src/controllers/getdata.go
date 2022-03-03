package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/paij0se/ymp3cli-api/src/database"
	"github.com/paij0se/ymp3cli-api/src/interfaces"
)

func GetData(ctx *gin.Context) {
	var (
		ymp3cli *interfaces.Ymp3cli
		id      *int
	)

	db, err := database.Connect()

	if err != nil {
		log.Println(err.Error())

		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "500 - Internal Server Error.",
			"error":   err.Error(),
		})

		return
	}

	if err = database.Query(db, id, ymp3cli); err != nil {
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
		"id": id,

		"app":      ymp3cli.App,
		"username": ymp3cli.Username,
	})
}

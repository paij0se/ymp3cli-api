package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/paij0se/ymp3cli-api/src/database"
	"github.com/paij0se/ymp3cli-api/src/interfaces"
)

func GetData(ctx *gin.Context) {
	var API []interfaces.Ymp3cli = []interfaces.Ymp3cli{}

	db, err := database.Connect()

	if err != nil {
		log.Println(err.Error())

		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "500 - Internal Server Error.",
			"error":   err.Error(),
		})

		return
	}

	if err = database.Query(db, &API); err != nil {
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

	ctx.JSON(200, API)
}

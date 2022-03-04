package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/paij0se/ymp3cli-api/src/database"
	"github.com/paij0se/ymp3cli-api/src/interfaces"
)

type ma2x struct {
	Max int
}

func GetData(ctx *gin.Context) {
	var (
		API []interfaces.Ymp3cli = []interfaces.Ymp3cli{}
		max ma2x
	)

	reqBody, _ := ioutil.ReadAll(ctx.Request.Body)

	if err := json.Unmarshal(reqBody, &max); err != nil {
		max = ma2x{Max: 20}
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

	if err = database.Query(db, &API, uint64(max.Max)); err != nil {
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

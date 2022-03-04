package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/paij0se/ymp3cli-api/src/controllers"
	"github.com/paij0se/ymp3cli-api/src/middlewares"
)

func SetupRouter(port string) error {
	router := gin.New()

	// Middlewares
	router.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		c.Next()
	})

	// Routes
	router.GET("/", middlewares.RateLimit("GET:/", 500), controllers.GetData)

	router.POST("/", middlewares.RateLimit("POST:/", 20000), controllers.UpdateData)

	return router.Run(":" + port)
}

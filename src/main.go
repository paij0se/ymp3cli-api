package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/paij0se/ymp3cli-api/src/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")

	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Println("The port to use is not declared, using port 8080.")

		port = "8080"
	}

	log.Println("Serve on port " + port)
	log.Fatalln(routes.SetupRouter(port))
}

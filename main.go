package main

import (
	"basic-trade/database"
	"basic-trade/routers"
	"os"
)

func main() {
	var PORT = os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080" // Default port if PORT environment variable is not set
	}

	database.StartDB()

	routers.StartServer().Run(":" + PORT)
}

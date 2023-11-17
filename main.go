package main

import (
	"assignment-project-rest-api/database"
	"assignment-project-rest-api/routers"
)

func main() {
	var PORT = ":8080"

	database.StartDB()

	routers.StartServer().Run(PORT)
}

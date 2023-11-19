package main

import (
	"basic-trade/database"
	"basic-trade/routers"
)

func main() {
	var PORT = ":8080"

	database.StartDB()

	routers.StartServer().Run(PORT)
}

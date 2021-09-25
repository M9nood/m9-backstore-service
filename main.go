package main

import (
	"m9-backstore-service/database"
	"os"

	router "m9-backstore-service/routes"
)

func main() {
	db := database.DBConnect(os.Getenv("DATABASE_URL"))
	router.RouterSetup()
	defer db.Close()
}

package main

import (
	"m9-backstore-service/database"
	router "m9-backstore-service/routes"
	"os"
)

func main() {
	db := database.DBConnect(os.Getenv("DATABASE_URL"))
	router.RouterSetup()
	defer db.Close()
}

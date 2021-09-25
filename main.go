package main

import (
	router "m9-backstore-service/routes"
)

func main() {
	// db := database.DBConnect(os.Getenv("DB_CONN"))
	router.RouterSetup()
	// defer db.Close()
}

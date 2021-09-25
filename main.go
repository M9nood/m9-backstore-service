package main

import (
	"log"
	"m9-backstore-service/database"
	"os"

	router "m9-backstore-service/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := database.DBConnect(os.Getenv("DATABASE_URL"))
	router.RouterSetup()
	defer db.Close()
}

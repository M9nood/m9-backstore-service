package main

import (
	"log"
	"m9-backstore-service/database"
	"os"

	router "m9-backstore-service/routes"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db := database.DBConnect(os.Getenv("DATABASE_URL"))
	router.RouterSetup()
	defer db.Close()
}

package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

func Connect() *gorm.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbName, port)

	Db, err := gorm.Open("postgres", dsn)

	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}
	fmt.Println("database connected.")
	return Db
}

func GetDB() *gorm.DB {
	return Db
}

func CloseDB() {
	Db.Close()
}

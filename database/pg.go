package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

var Db *pg.DB

func DBConnect(url string) *pg.DB {
	url = "postgres://vqlrlplcqgtkot:4d00ea2ad1bc6d128b1d8ad64b31084c004f266487575e568b85c815e3590de6@ec2-3-220-214-162.compute-1.amazonaws.com:5432/dfv0fqufuui5sq"
	opt, err := pg.ParseURL(url)
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	db := pg.Connect(opt)
	return db
}

func GetDB() *pg.DB {
	return Db
}

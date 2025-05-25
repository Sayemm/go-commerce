package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() {
	var err error
	DB, err = sqlx.Connect("mysql", "root:password@tcp(product-db:3306)/productdb")
	if err != nil {
		log.Fatalln("Database Connection Failed: ", err)
	}
}

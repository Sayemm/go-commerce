package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Connect() {
	var err error

	dsn := "root:pass@tcp(127.0.0.1:3306)/ecommerce"

	DB, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("Failed to connect to the database: ", err)
	}
	fmt.Println("Connected to the database!s")
}

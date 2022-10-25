package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)




func  Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "root:root@(localhost:3306)/bank_go?parseTime=true")

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}


	return db,nil

}


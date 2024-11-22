package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	var err error
	DB, err := sql.Open("mysql", "test_user:test123@tcp(156.67.219.200:3306)/hoopla_db")
	if err != nil {
		log.Print("Error connecting to the database: ", err)
		log.Fatal(err)
	}

	// Check the connection
	if err = DB.Ping(); err != nil {
		log.Print("Error pinging the database: ", err)
		log.Fatal(err)
	}

	log.Print("Connected to the database")
	return DB
}

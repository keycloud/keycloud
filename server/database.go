package main


import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


// debug purpose
const (
	host     = "localhost"
	port     = 5432
	user     = "keycloud"
	password = "keycloud"
	dbname   = "postgres"
)


func connectDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// test database connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
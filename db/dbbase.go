package db

import (
	"database/sql"
	"fmt"
)

// DbOperator interface
type DbOperator interface {
	Insert(interface{}) error
}

const (
	host     = "192.168.31.163"
	port     = 5432
	user     = "postgres"
	password = "w184485789"
	dbname   = "ocrdb"
)

var db *sql.DB

func connectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

// Default db
func Default() *sql.DB {
	if db == nil {
		db = connectDB()
	}

	return db
}

package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Db struct {
	ConnectionStr string
	Db            *sql.DB
}

var db Db

func Open(ConnectionStr string) {
	var err error

	db.ConnectionStr = ConnectionStr
	db.Db, err = sql.Open("postgres", db.ConnectionStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Db.Ping(); err != nil {
		log.Fatal(err)
	}

	db.Db.SetMaxIdleConns(10)
	db.Db.SetMaxOpenConns(20)
	db.Db.SetConnMaxLifetime(time.Hour)

	log.Print("Connected to database.")
}

func Close() {
	if db.Db != nil {
		if err := db.Db.Close(); err != nil {
			log.Fatal(err)
		}
		db.Db = nil
		log.Print("Succesfully closed the connection.")
		return
	}
	log.Print("Cannot close unopened connection.")
}

func Connect() (*sql.Conn, error) {
	return db.Db.Conn(context.Background())
}

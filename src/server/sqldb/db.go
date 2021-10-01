package sqldb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB 

func ConnectDB() *sql.DB {
    db,err := sql.Open("sqlite3", "./client.db")
	if err != nil {
		panic(err.Error())
	}

	DB = db

	return db
}
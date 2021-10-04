/* Simply sets up a connection to the database and returns a global variable that can be used in any package */

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
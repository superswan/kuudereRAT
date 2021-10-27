/*
This package deals with registering and updating clients
*/
package client

import (
	"log"
	"database/sql"
	"desukit/sqldb"
	"desukit/sqldb/model"
)

// Check for existince in DB
func CheckClient(client model.Client) (bool) {
	var exists bool
	if err := sqldb.DB.QueryRow("SELECT id FROM clients WHERE uuid=$1", 
		client.Uuid).Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	return exists 
}

func GetClient(uuid string) (model.Client) {
	C := new(model.Client)
	err := sqldb.DB.QueryRow("SELECT uuid, ipaddr, ts_first, ts_last FROM clients where uuid =?",
	uuid).Scan(&C.Uuid, &C.Ipaddr, &C.Ts_first, &C.Ts_last)
	
	if err != nil {
		log.Fatal(err)
	}
	if err == sql.ErrNoRows {
		log.Fatal(err)
	}

	return *C
}

// Insert
// Register client into database
func RegisterClient(client model.Client) {
	statement, err := sqldb.DB.Prepare("INSERT INTO clients (uuid, ipaddr, platform, ts_first, ts_last) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = statement.Exec(client.Uuid, client.Ipaddr, client.Platform, client.Ts_first, client.Ts_last)
	if err != nil {
		log.Fatal(err)
	}
}

// Update
// Update Timestamp
func UpdateClient(client model.Client) {
	statement, err := sqldb.DB.Prepare("UPDATE clients SET ts_last = ? WHERE uuid = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(client.Ts_last, client.Uuid)
	if err != nil {
		log.Fatal(err)
	}
}

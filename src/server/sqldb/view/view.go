package view

import (
	"desukit/sqldb"
	"desukit/sqldb/model"
)

func RetrieveClients() []model.Client {
	rows, _ := sqldb.DB.Query("SELECT uuid, ipaddr, ts_last from clients")
	
	Clients := []model.Client{}

	defer rows.Close()
	for rows.Next() {
		C := new(model.Client)
		rows.Scan(&C.Uuid, &C.Ipaddr, &C.Ts_last)
		Clients = append(Clients, *C)
	}

	return Clients
}
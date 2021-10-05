package schema

import (
	"desukit/sqldb"
)

func InitializeDB() {
    // table: clients
    // columns:
    // id primary key
    // uuid text not null unique
    // ipaddr text
    // ts_first text not null -- timestamp of when first seen
    // ts_last text not null -- most recent timestamp
    statement, _ := sqldb.DB.Prepare("CREATE TABLE IF NOT EXISTS clients (id INTEGER PRIMARY KEY, uuid TEXT NOT NULL UNIQUE, ipaddr TEXT, ts_first TEXT NOT NULL, ts_last TEXT NOT NULL)")
    statement.Exec() 

    // table: tasks
    // columns:
    // id primary key
    // uuid text not null unique
    // task_queued integer
    // queue int
    // ts_last text -- timestamp of last execution
    statement, _ = sqldb.DB.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, uuid TEXT NOT NULL UNIQUE, task_id INTEGER, task_queued INTEGER, queue INTEGER, ts_last TEXT)")
    statement.Exec()
    // ------------ END DB ----------------------
}
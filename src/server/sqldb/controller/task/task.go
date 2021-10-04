/*
All task management for clients. 
*/
package tasks

import (
	"log"
	"database/sql"
	"desukit/sqldb"
	"desukit/sqldb/model"
)

// Task checking
func CheckTasks(client model.Client) {
	var task_queued int
	task_row := sqldb.DB.QueryRow("SELECT task_queued FROM tasks WHERE uuid = $1", 
		client.Uuid)
	task_row.Scan(&task_queued)

	return task_queued
}

func CreateTask(client model.Client, task model.Task) {

}

func ManageTaskQueue(client model.Client) {
	statement, _ := sqldb.DB.Prepare("UPDATE tasks SET task_queued = 0, ts_last = ? WHERE uuid = ?")
	statement.Exec(timestamp, uuid)
}
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

// Check task existence 
func CheckTasks(client model.Client) (bool) {
	var task_exists bool
	if err := sqldb.DB.QueryRow("SELECT id FROM tasks WHERE uuid = $1", 
		client.Uuid).Scan(&task_exists); err != nil && err != sql.ErrNoRows {
			log.Fatal(err)
	}
	return task_exists;	
}

func GetTask(client model.Client) (task model.Task) {
	T := new(model.Task)
	err := sqldb.DB.QueryRow("SELECT uuid, task_id, task_queued, queue, ts_last FROM tasks where uuid =?", 
	client.Uuid).Scan(&T.Uuid, &T.Task_ID, &T.Task_queued, &T.Queue, &T.Ts_last)
	
	if err != nil {
		log.Fatal(err)
	}
	if err == sql.ErrNoRows {
		log.Fatal(err)
	}

	return *T
}

//Only using a single queue at this time
func CreateTask(client model.Client, task model.Task) (bool) {
	task_exists := CheckTasks(client)
	if task_exists {
		return false
	}
	// Just realized task_queued value is completely unneccessary at this time. Will keep for
	// future use  
	task_queued := 1
	queue := 1

	statement, _ := sqldb.DB.Prepare("INSERT INTO tasks (uuid, task_id, task_queued, queue, ts_last) VALUES (?, ?, ?, ?, ?)")
	statement.Exec(client.Uuid, task.Task_ID, task_queued, queue, client.Ts_last)

	return true
}

func ClearTaskQueue(client model.Client) {
	statement, err := sqldb.DB.Prepare("DELETE FROM tasks WHERE uuid = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(client.Uuid)
	if err != nil {
		log.Fatal(err)
	}
}
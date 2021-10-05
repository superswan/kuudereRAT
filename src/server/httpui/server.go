package httpui

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"

	"desukit/sqldb/view"
	"desukit/sqldb/model"
	"desukit/sqldb/controller/client"
	"desukit/sqldb/controller/tasks"
)

var templates *template.Template

func Start() {
	port := "8080"
	
	templates = template.Must(template.ParseGlob("httpui/templates/*.html"))
	r := mux.NewRouter()	
	r.HandleFunc("/", httpHandler)
	http.Handle("/", r)

	// Task functions
	r.HandleFunc("/create_task/{uuid}", taskHandler)

	// Static files
	fileServer := http.FileServer(http.Dir("./httpui/static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	// Generated payloads will be placed in this directory and can be used to
	// push modules and client payloads to remote hosts
	pullServer := http.FileServer(http.Dir("./httpui/payload"))
	r.PathPrefix("/checkout/").Handler(http.StripPrefix("/checkout/", pullServer))

	done := make(chan bool)
	fmt.Println("Starting HTTP server at port 8080")
	http.ListenAndServe(":"+port, nil)
	<-done
}
// Helper functions
func createTask(uuid string, task_name string) {
	var task_id int
	switch task_name {
	case "shell":
		task_id = 1
	}
	// fill Task struct with garbage, won't actually be submitted right now
	task_queued := 1
	queue := 1
	ts_last := "timestamp"

	C := client.GetClient(uuid)
	T := model.Task{uuid, task_id, task_queued, queue, ts_last}

	tasks.CreateTask(C, T)
}

// Handlers

func httpHandler(w http.ResponseWriter, r *http.Request) {
	Clients := view.RetrieveClients()
	err := templates.ExecuteTemplate(w, "index.html", Clients)
	if err != nil {
		log.Fatal("Error in template: ", err)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	r.ParseForm()
	task_name := r.Form["task"][0]
	createTask(uuid, task_name)
	httpHandler(w, r)
}
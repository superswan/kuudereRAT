package httpui

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"

	"b00fkit/sqldb"
)

var templates *template.Template

type Client struct {
	Id int
	Uuid string
	Ipaddr string
	Timestamp string
}

func Start() {
	port := "8080"
	
	templates = template.Must(template.ParseGlob("httpui/templates/*.html"))
	r := mux.NewRouter()	
	r.HandleFunc("/", httpHandler)
	http.Handle("/", r)
	fileServer := http.FileServer(http.Dir("./httpui/static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
	// Generated payloads will be placed in this directory and can be used to
	// push modules and payloads to remote hosts
	pullServer := http.FileServer(http.Dir("./httpui/payload"))
	r.PathPrefix("/checkout/").Handler(http.StripPrefix("/checkout/", pullServer))

	done := make(chan bool)
	fmt.Println("Starting HTTP server at port 8080")
	http.ListenAndServe(":"+port, nil)
	<-done
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	rows, _ := sqldb.DB.Query("SELECT * from clients")
	
	Clients := []Client{}

	defer rows.Close()
	for rows.Next() {
		C := new(Client)
		rows.Scan(&C.Id, &C.Uuid, &C.Ipaddr, &C.Timestamp)
		Clients = append(Clients, *C)
	}

	templates.ExecuteTemplate(w, "index.html", Clients)
}
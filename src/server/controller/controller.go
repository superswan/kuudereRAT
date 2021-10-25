//Optional module API, incomplete

package main

import (
	"fmt"
	"log"
	"net/http"
	//"encoding/json"
	//"encoding/base64"
	"github.com/gorilla/mux"
)

type Host struct (
	Uuid		string `json:"uuid"`
	Os_info		string `json:"os"`
	Location	string `json:"country"`
	User		string `json:"user"`
)

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/register", register).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func main() {
	fmt.Println("It works!")
}

func register(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var H Host
	json.Unmarshal(reqBody, &H)
	
	fmt.Println(H)
}
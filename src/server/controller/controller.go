package controller

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"encoding/base64"
	"github.com/gorilla/mux"
)

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/register", register)

	log.Fatal(http.ListenAndServe(":3000", r))
}

func main() {
	fmt.Println("It works!")
}
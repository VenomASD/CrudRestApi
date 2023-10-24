package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
// HandlerRouting deeclares and sets up API end points
func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/data", GetData).Methods("GET")
	r.HandleFunc("/data", CreateData).Methods("POST")
	r.HandleFunc("/data", UpdateData).Methods("PUT")
	r.HandleFunc("/data", DeleteData).Methods("DELETE")

	log.Println("Listening on port 8080..........")
	http.ListenAndServe(":8080", r)
}

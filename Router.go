package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/data", GetData).Methods("GET")
	r.HandleFunc("/data", CreateData).Methods("POST")
	r.HandleFunc("/data", UpdateData).Methods("PUT")
	r.HandleFunc("/data", DeleteData).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}

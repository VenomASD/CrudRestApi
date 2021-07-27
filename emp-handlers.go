package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)


func CreateEmployee(w http.ResponseWriter , r *http.Request) {

	w.Header().Set("Content-Type","application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)

}

func GetEmployees(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var emp []Employee
	Database.Find(&emp)
	json.NewEncoder(w).Encode(emp)	
}

func GetEmployeeByID(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var emp Employee
	Database.First(&emp , mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(emp)		
}

func UpdateEmployee(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var emp Employee
	Database.First(&emp , mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Save(&emp)
	json.NewEncoder(w).Encode(emp)		
}

func DeleteEmployee(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var emp Employee
	Database.Delete(&emp , mux.Vars(r)["eid"])

	json.NewEncoder(w).Encode("employee is deleted")		
}


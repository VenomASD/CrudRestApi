package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func CreateData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var act Actor
	json.NewDecoder(r.Body).Decode(&act)
	_, err := Database.Raw("Insert into actor (actor_id, first_name, last_name) values (?,?,?)", act.ActorId, act.FirstName, act.LastName).Rows()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode("Successfully created a record in db")
}

func GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var act []Actor
	id := r.URL.Query().Get("id")
	// Raw SQL Query
	rows, err := Database.Raw("select * from actor where actor_id = ?", id).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var newActor Actor
		rows.Scan(
			&newActor.ActorId,
			&newActor.FirstName,
			&newActor.LastName,
			&newActor.TimeStamp,
		)
		act = append(act, newActor)
		// do something
	}
	json.NewEncoder(w).Encode(act)
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var act Actor
	json.NewDecoder(r.Body).Decode(&act)
	_, err := Database.Raw("Update actor set first_name=?, last_name=? where actor_id=?", act.FirstName, act.LastName, act.ActorId).Rows()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode("Successfully updated a record in db")
}

func DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	_, err := Database.Raw("Delete from actor where actor_id=?", id).Rows()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode("Successfully deleted a record in db")
}

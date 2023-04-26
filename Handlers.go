package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func CreateData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var act Actor
	json.NewDecoder(r.Body).Decode(&act)
	QueryTemplate := fmt.Sprintf(CreateDataQuery, strconv.Itoa(act.ActorId), act.FirstName, act.LastName)
	_, err := Database.Raw(QueryTemplate).Rows()
	log.Println(QueryTemplate)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("Successfully created a record in db")
}

func GetData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var act []Actor
	id := r.URL.Query().Get("id")
	// Raw SQL Query
	QueryTemplate := fmt.Sprintf(GetDataQuery, id)
	rows, err := Database.Raw(QueryTemplate).Rows()
	log.Println(QueryTemplate)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
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
	}
	json.NewEncoder(w).Encode(act)
}
func GetAllData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var act []Actor
	// Raw SQL Query
	QueryTemplate := fmt.Sprintf(GetAllDataQuery)
	rows, err := Database.Raw(QueryTemplate).Rows()
	log.Println(QueryTemplate)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
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
	}
	json.NewEncoder(w).Encode(act)
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var act Actor
	json.NewDecoder(r.Body).Decode(&act)
	QueryTemplate := fmt.Sprintf(UpdateDataQuery, act.FirstName, act.LastName, strconv.Itoa(act.ActorId))
	_, err := Database.Raw(QueryTemplate).Rows()
	log.Println(QueryTemplate)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("Successfully updated a record in db")
}

func DeleteData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	QueryTemplate := fmt.Sprintf(DeleteDataQuery , id)
	_, err := Database.Raw(QueryTemplate).Rows()
	log.Println(QueryTemplate)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("Successfully deleted a record in db")
}

func ProduceData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var msg string
	json.NewDecoder(r.Body).Decode(msg)
	log.Println(msg)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	StartProducer(msg)
	json.NewEncoder(w).Encode("Success processed req!")
}
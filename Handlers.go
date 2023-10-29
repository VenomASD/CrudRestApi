package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func CreateData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var act Actor
	json.NewDecoder(r.Body).Decode(&act)
	QueryTemplate := fmt.Sprintf(CreateDataQuery, strconv.Itoa(act.ActorId), act.FirstName, act.LastName)
	_, err := Database.Raw(QueryTemplate).Rows()
	log.Println(QueryTemplate)
	if err != nil {
		log.Println(err)
		ServiceJSONResponse(w, http.StatusText(500), "", err)
		return
	}
	ServiceJSONResponse(w, http.StatusText(200), RecordCreatedSuccess, "")
}

func GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var act []*Actor
	id := r.URL.Query().Get("id")
	// Raw SQL Query
	QueryTemplate := fmt.Sprintf(GetDataQuery, id)
	rows, err := Database.Raw(QueryTemplate).Rows()
	log.Println(QueryTemplate)
	if err != nil {
		log.Println(err)
		ServiceJSONResponse(w, http.StatusText(500), "", err)
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
		act = append(act, &newActor)
	}
	// no record found in db
	if len(act) == 0 {
		ServiceJSONResponse(w, http.StatusText(404), RecordNotFound, "")
		return
	}
	ServiceJSONResponse(w, http.StatusText(200), "", act)
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var act Actor
	json.NewDecoder(r.Body).Decode(&act)
	QueryTemplate := fmt.Sprintf(UpdateDataQuery, act.FirstName, act.LastName, strconv.Itoa(act.ActorId))
	_, err := Database.Raw(QueryTemplate).Rows()
	log.Println(QueryTemplate)
	if err != nil {
		log.Println(err)
		ServiceJSONResponse(w, http.StatusText(500), "", err)
		return
	}
	ServiceJSONResponse(w, http.StatusText(200), RecordUpdatedSuccess, "")
}

func DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	QueryTemplate := fmt.Sprintf(DeleteDataQuery, id)
	_, err := Database.Raw(QueryTemplate).Rows()
	log.Println(QueryTemplate)
	if err != nil {
		log.Println(err)
		ServiceJSONResponse(w, http.StatusText(500), "", err)
		return
	}
	ServiceJSONResponse(w, http.StatusText(200), RecordDeletedSuccess, "")
}

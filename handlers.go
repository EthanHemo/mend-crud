package main

import (
	"encoding/json"
	"mend-crud/databases"
	"net/http"
)

type Payload struct {
	Name      string
	Superhero databases.Superhero
}

var db databases.Crud

func AddSuperheroHandler(w http.ResponseWriter, r *http.Request) {
	var s databases.Superhero
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = db.AddSuperhero(s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RemoveSuperheroHandler(w http.ResponseWriter, r *http.Request) {
	var p Payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = db.DeleteSuperhero(p.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateSuperheroHandler(w http.ResponseWriter, r *http.Request) {
	var p Payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = db.UpdateSuperhero(p.Name, p.Superhero); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetSuperheroHandler(w http.ResponseWriter, r *http.Request) {
	var p Payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s, err := db.GetSuperhero(p.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

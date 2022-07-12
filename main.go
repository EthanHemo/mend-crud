package main

import (
	"mend-crud/databases"
	"net/http"
)

func main() {
	db = databases.GetDatabase(databases.Array)

	http.HandleFunc("/add", AddSuperheroHandler)
	http.HandleFunc("/remove", RemoveSuperheroHandler)
	http.HandleFunc("/update", UpdateSuperheroHandler)
	http.HandleFunc("/get", GetSuperheroHandler)

	if err := http.ListenAndServeTLS("localhost:9443", "go-server1.crt", "go-server1.key", nil); err != nil {
		panic(err)
	}
}

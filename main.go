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

	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
	// if err := http.ListenAndServeTLS("localhost:9443", "server.crt", "server.key", nil); err != nil {
	// 	panic(err)
	// }
}

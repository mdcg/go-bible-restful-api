package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdcg/go-bible-restful-api/api/routes"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v0").Subrouter()

	// Books
	api.HandleFunc("/books", routes.FindAllBooks).Methods(http.MethodGet)

	// Testaments
	api.HandleFunc("/testaments", routes.FindAllTestaments).Methods(http.MethodGet)

	// Verses
	api.HandleFunc("/verses", routes.FindAllVerses).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}

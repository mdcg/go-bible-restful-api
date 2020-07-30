package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdcg/go-bible-restful-api/api/routes"
)

func main() {
	// conn := db.GetDB()
	// defer conn.Close()

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v0").Subrouter()
	api.HandleFunc("/", routes.SanityTest).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}

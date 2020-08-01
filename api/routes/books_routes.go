package routes

import (
	"net/http"

	"github.com/mdcg/go-bible-restful-api/api/db"

	"github.com/gorilla/mux"
	"github.com/mdcg/go-bible-restful-api/api/utils"

	"github.com/mdcg/go-bible-restful-api/api/controllers"
)

func FindAllBooks(w http.ResponseWriter, r *http.Request) {
	books := controllers.FindAllBooks()
	utils.JSONResponse(w, http.StatusFound, books)
}

func FindBookByAbbrev(w http.ResponseWriter, r *http.Request) {
	abbrev := mux.Vars(r)["abbrev"]
	book := controllers.FindBookByAbbrev(abbrev)

	if utils.WasAnyDataFound(book, db.Books{}) {
		utils.JSONResponse(w, http.StatusFound, book)
	}
	utils.JSONResponse(w, http.StatusNotFound, map[string]string{})
}

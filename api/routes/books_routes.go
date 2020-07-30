package routes

import (
	"net/http"

	"github.com/mdcg/go-bible-restful-api/api/utils"

	"github.com/mdcg/go-bible-restful-api/api/controllers"
)

func FindAllBooks(w http.ResponseWriter, r *http.Request) {
	books := controllers.FindAllBooks()
	utils.JSONResponse(w, 200, books)
}

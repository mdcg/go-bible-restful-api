package routes

import (
	"net/http"

	"github.com/mdcg/go-bible-restful-api/api/utils"

	"github.com/mdcg/go-bible-restful-api/api/controllers"
)

func FindAllVerses(w http.ResponseWriter, r *http.Request) {
	verses := controllers.FindAllVerses()
	utils.JSONResponse(w, 200, verses)
}

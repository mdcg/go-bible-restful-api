package routes

import (
	"net/http"

	"github.com/mdcg/go-bible-restful-api/api/utils"

	"github.com/mdcg/go-bible-restful-api/api/controllers"
)

func FindAllTestaments(w http.ResponseWriter, r *http.Request) {
	testaments := controllers.FindAllTestaments()
	utils.JSONResponse(w, http.StatusFound, testaments)
}

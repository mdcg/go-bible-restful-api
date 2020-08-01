package routes

import (
	"net/http"

	"github.com/mdcg/go-bible-restful-api/api/utils"

	"github.com/mdcg/go-bible-restful-api/api/controllers"
)

func FindAllVersions(w http.ResponseWriter, r *http.Request) {
	versions := controllers.FindAllVersions()
	utils.JSONResponse(w, 200, versions)
}

package routes

import (
	"fmt"
	"net/http"
	"os"
)

func SanityTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(os.Getenv("MYSQL_PASSWORD"))
	w.Write([]byte(`{"message": "get called"}`))
}

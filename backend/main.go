package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	var port = os.Getenv("BACKEND_PORT")
	if len(port) == 0 {
		port = "8000"
	}
	router := mux.NewRouter()
	CreateAPIRouter(router)
	LoadMiddleware(router)
	http.ListenAndServe(":"+port, router)
}

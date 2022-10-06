package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CreateAPIRouter(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()
	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	api.HandleFunc("/blog", SingleBlogHandler).Methods(http.MethodGet)        //, http.MethodOptions)
	api.HandleFunc("/blog/{slug}", SingleBlogHandler).Methods(http.MethodGet) //, http.MethodOptions)
}

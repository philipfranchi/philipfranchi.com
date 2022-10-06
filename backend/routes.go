package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CreateAPIRouter(handler *APIHandler) *mux.Router {
	router := mux.NewRouter()
	LoadMiddleware(router)
	api := router.PathPrefix("/api").Subrouter()
	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	api.HandleFunc("/blog", handler.SingleBlogHandler).Methods(http.MethodGet)        //, http.MethodOptions)
	api.HandleFunc("/blog/{slug}", handler.SingleBlogHandler).Methods(http.MethodGet) //, http.MethodOptions)
	return router
}

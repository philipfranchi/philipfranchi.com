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
	AddBlogSubRouter(api, handler)
	return router
}

func AddBlogSubRouter(router *mux.Router, handler *APIHandler) {
	blog := router.PathPrefix("/blog").Subrouter()
	blog.HandleFunc("/", handler.HandleGetAllBlogPostMetadata).Methods(http.MethodGet)
	blog.HandleFunc("/{slug}", handler.HandleGetSingleBlogPost).Methods(http.MethodGet)
}

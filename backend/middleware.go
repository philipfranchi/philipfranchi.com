package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
func LoadMiddleware(router *mux.Router) {
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(loggingMiddleware)
}

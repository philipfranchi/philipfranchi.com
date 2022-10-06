package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func respond(w http.ResponseWriter, code int, msg []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}

func respondWithError(w http.ResponseWriter, err *ApplicationError) {
	respond(w, err.Code, []byte(err.Message))
}

func SingleBlogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	post, appErr := GetBlogPostBySlug(slug)
	if appErr != nil {
		respondWithError(w, appErr)
	}
	data, err := json.Marshal(post)
	if err != nil {
		respondWithError(w, MarshallingError(err.Error()))
	}
	respond(w, http.StatusAccepted, data)
}

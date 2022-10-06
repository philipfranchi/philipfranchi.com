package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Respond(w http.ResponseWriter, code int, msg []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}

func RespondWithError(w http.ResponseWriter, err *ApplicationError) {
	Respond(w, err.Code, []byte(err.Message))
}

func SingleBlogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	post, appErr := GetBlogPostBySlug(slug)
	if appErr != nil {
		RespondWithError(w, appErr)
	}
	data, err := json.Marshal(post)
	if err != nil {
		RespondWithError(w, MarshallingError(err.Error()))
	}
	Respond(w, http.StatusAccepted, data)
}

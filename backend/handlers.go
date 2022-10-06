package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type APIHandler struct {
	blog *BlogProvider
}

func CreateAPIHandler(blog *BlogProvider) *APIHandler {
	handler := APIHandler{blog}
	return &handler
}

func (h *APIHandler) respond(w http.ResponseWriter, code int, msg []byte) {
	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}

func (h *APIHandler) respondWithError(w http.ResponseWriter, err *ApplicationError) {
	h.respond(w, err.Code, []byte(err.Message))
}

func (h *APIHandler) SingleBlogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	post, appErr := h.blog.GetBlogPostBySlug(slug)
	if appErr != nil {
		h.respondWithError(w, appErr)
	}
	data, err := json.Marshal(post)
	if err != nil {
		h.respondWithError(w, MarshallingError(err.Error()))
	}
	h.respond(w, http.StatusAccepted, data)
}

package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type APIHandler struct {
	blog BlogManager
}

func CreateAPIHandler(blog BlogManager) *APIHandler {
	handler := APIHandler{blog}
	return &handler
}

func (h *APIHandler) respond(w http.ResponseWriter, code int, msg []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(msg)
}

func (h *APIHandler) respondWithError(w http.ResponseWriter, err *ApplicationError) {
	h.respond(w, err.Code, []byte(err.Message))
}

func (h *APIHandler) HandleGetAllBlogPostMetadata(w http.ResponseWriter, r *http.Request) {
	posts, appErr := h.blog.GetAllBlogPostMetadata()
	if appErr != nil {
		h.respondWithError(w, appErr)
		return
	}

	payload, err := json.Marshal(posts)
	if err != nil {
		h.respondWithError(w, MarshallingError(err.Error()))
		return
	}
	h.respond(w, http.StatusOK, payload)
}

func (h *APIHandler) HandleGetSingleBlogPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	if len(slug) == 0 {
		h.respondWithError(w, ValidationError())
		return
	}

	post, appErr := h.blog.GetBlogPostBySlug(slug)
	if appErr != nil {
		h.respondWithError(w, appErr)
		return
	}
	data, err := json.Marshal(post)
	if err != nil {
		h.respondWithError(w, MarshallingError(err.Error()))
		return
	}
	h.respond(w, http.StatusOK, data)
}

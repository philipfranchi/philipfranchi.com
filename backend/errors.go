package main

import (
	"fmt"
	"net/http"
)

type ApplicationError struct {
	Code    int
	Message string
}

func (err *ApplicationError) Error() string {
	return err.Message
}

func ValidationError() *ApplicationError {
	err := ApplicationError{
		http.StatusBadRequest, "bad slug",
	}
	return &err
}

func PostMissingError(slug string) *ApplicationError {
	err := ApplicationError{
		http.StatusNotFound, fmt.Sprintf("Post with id: %s not found", slug),
	}
	return &err
}

func BlogParseError(msg string) *ApplicationError {
	err := ApplicationError{
		http.StatusInternalServerError, "Issue parsing blog: " + msg,
	}
	return &err
}

func MarshallingError(msg string) *ApplicationError {
	err := ApplicationError{
		http.StatusInternalServerError, msg,
	}
	return &err
}

func InternalServerError(msg string) *ApplicationError {
	err := ApplicationError{
		http.StatusInternalServerError, msg,
	}
	return &err
}

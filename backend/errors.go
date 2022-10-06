package main

import "net/http"

type ApplicationError struct {
	Code    int
	Message string
}

func (err *ApplicationError) Error() string {
	return err.Message
}

func ValidationError(msg string) *ApplicationError {
	err := ApplicationError{
		http.StatusBadRequest, msg,
	}
	return &err
}

func MarshallingError(msg string) *ApplicationError {
	err := ApplicationError{
		http.StatusInternalServerError, msg,
	}
	return &err
}

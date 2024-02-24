package api

import (
	"encoding/json"
	"net/http"
)

type SearchParams struct {
	Username string
}

type Balance struct {
	StatusCode int
	Balance    int
}

type Error struct {
	ErrorCode int
	Message   string
}

func WriteError(w http.ResponseWriter, message string, code int) {
	response := Error{
		ErrorCode: code,
		Message:   message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		WriteError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		WriteError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)

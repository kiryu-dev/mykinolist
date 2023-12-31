package controller

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func writeJSONResponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func writeErrorJSON(w http.ResponseWriter, status int, errMessage string) error {
	errResponse := &errorResponse{Error: errMessage}
	return writeJSONResponse(w, status, errResponse)
}

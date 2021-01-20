package rest

import (
	"encoding/json"
	"net/http"
)

type successResponse struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func responseSuccess(w http.ResponseWriter, data interface{}) {
	payload, convertErr := json.Marshal(successResponse{data})
	if convertErr != nil {
		responseError(w, http.StatusInternalServerError, convertErr)
		return
	}
	response(w, http.StatusOK, payload)
}

func responseError(w http.ResponseWriter, errorStatusCode int, err error) {
	payload, convertErr := json.Marshal(errorResponse{err.Error()})
	if convertErr != nil {
		responseError(w, http.StatusInternalServerError, convertErr)
		return
	}
	response(w, errorStatusCode, payload)
}

func response(w http.ResponseWriter, statusCode int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(payload)
}

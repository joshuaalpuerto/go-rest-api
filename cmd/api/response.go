package main

import (
	"encoding/json"
	"net/http"
)

// Response represents a standardized API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

// JSONResponse sends a JSON response with the given status code
func JSONResponse(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(Response{
		Success: statusCode < 400,
		Data:    data,
	})

	return err
}

// JSONError sends a JSON error response
func JSONError(w http.ResponseWriter, statusCode int, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(Response{
		Success: false,
		Error:   message,
	})

	return err
}

// SuccessResponse sends a success response with data
func SuccessResponse(w http.ResponseWriter, data any) error {
	return JSONResponse(w, http.StatusOK, data)
}

// ErrorResponse sends an error response
func ErrorResponse(w http.ResponseWriter, statusCode int, message string) error {
	return JSONError(w, statusCode, message)
}

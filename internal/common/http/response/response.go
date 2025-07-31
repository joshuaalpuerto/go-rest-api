package response

import (
	"encoding/json"
	"net/http"
)

// Response represents a standardized API response
type Response struct {
	Data    any    `json:"data,omitempty"`
	Details string `json:"details,omitempty"`
	Status  int    `json:"status"`
	// TODO add metadata here for pagination later
}

func (r *Response) SendSuccessResponse(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	r.Data = data
	r.Status = statusCode

	if err := json.NewEncoder(w).Encode(r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (r *Response) SendErrorResponse(w http.ResponseWriter, details string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	r.Details = details
	r.Status = statusCode

	if err := json.NewEncoder(w).Encode(r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

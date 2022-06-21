package dto

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Ok bool `json:"ok"`
}

type Error struct {
	Status
	Message string `json:"err"`
}

func JSONError(w http.ResponseWriter, code int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(Error{
		Status:  Status{false},
		Message: err.Error(),
	})
}

func JSONResponse(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(data)
}

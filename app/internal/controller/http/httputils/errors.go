package httputils

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Status
	Message string `json:"err"`
}

func NewError(w http.ResponseWriter, code int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(Error{
		Status:  Status{false},
		Message: err.Error(),
	})
}

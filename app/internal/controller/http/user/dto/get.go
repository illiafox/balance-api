package dto

import (
	"fmt"
	"strconv"

	"balance-service/app/internal/controller/http/httputils"
)

type GetBalanceIN struct {
	UserID uint64 `json:"user_id"`
	Base   string `json:"base"`
}

func NewGetBalanceIN(id, base string) (GetBalanceIN, error) {
	var err error
	get := GetBalanceIN{
		Base: base,
	}
	// UserID
	if get.UserID, err = strconv.ParseUint(id, 10, 64); err != nil {
		return get, fmt.Errorf("parse id: %w", err)
	}
	if get.UserID <= 0 {
		return get, fmt.Errorf("id: expected > 0, got %d", get.UserID)
	}
	//
	return get, nil
}

// //

type Balance string

func (b Balance) MarshalJSON() ([]byte, error) {
	if len(b) > 0 {
		return []byte(b), nil
	}
	return []byte("null"), nil
}

type GetBalanceOUT struct {
	httputils.Status
	Balance Balance `json:"balance"`
	Base    string  `json:"base,omitempty"`
}

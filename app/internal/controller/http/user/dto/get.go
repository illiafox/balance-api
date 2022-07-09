package dto

import (
	"fmt"
)

type GetBalanceIN struct {
	UserID int64  `json:"user_id"`
	Base   string `json:"base"`
}

type Balance string

func (b Balance) MarshalJSON() ([]byte, error) {
	if len(b) > 0 {
		return []byte(b), nil
	}
	return []byte("null"), nil
}

func (g GetBalanceIN) Validate() error {
	if g.UserID <= 0 {
		return fmt.Errorf("invalid user id: got %d, expected > 0", g.UserID)
	}
	return nil
}

type GetBalanceOUT struct {
	Status
	Balance Balance `json:"balance"`
	Base    string  `json:"base,omitempty"`
}

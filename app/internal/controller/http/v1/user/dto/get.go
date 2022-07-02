package dto

import (
	"encoding/json"
	"fmt"
)

type GetBalanceIN struct {
	UserID int64  `json:"user_id"`
	Base   string `json:"base"`
}

func (g GetBalanceIN) Validate() error {
	if g.UserID <= 0 {
		return fmt.Errorf("invalid user id: got %d, expected > 0", g.UserID)
	}
	return nil
}

type GetBalanceOUT struct {
	Status
	Balance json.RawMessage `json:"balance"`
	Base    string          `json:"base,omitempty"`
}

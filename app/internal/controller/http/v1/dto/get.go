package dto

import "encoding/json"

type Status struct {
	Ok bool `json:"ok"`
}

type GetBalanceIn struct {
	UserID int64  `json:"user_id"`
	Base   string `json:"base"`
}

type GetBalanceOut struct {
	Status
	Balance json.RawMessage `json:"balance"`
	Base    string          `json:"base"`
}

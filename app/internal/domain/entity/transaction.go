package entity

import (
	"encoding/json"
	"time"
)

type Transaction struct {
	ID          int64           `json:"transaction_id"`
	ToID        int64           `json:"to_id"`   // to_id
	FromID      json.RawMessage `json:"from_id"` // number or "null"
	Action      int64           `json:"action"`
	Date        time.Time       `json:"date"`
	Description string          `json:"description"`
}

package entity

type Transaction struct {
	ID          int64  `json:"transaction_id"`
	ToID        int64  `json:"to_id"`             // to_id
	FromID      int64  `json:"from_id,omitempty"` // zero -> null -> received from other service
	Action      int64  `json:"action"`
	Date        Time   `json:"date"`
	Description string `json:"description"`
}

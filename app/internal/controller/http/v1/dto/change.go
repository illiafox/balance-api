package dto

type ChangeBalanceIn struct {
	UserID      int64  `json:"user_id"     validate:"required|min=1"`
	Amount      string `json:"base"        validate:"required"`
	Description string `json:"description" validate:"required|min=10"`
}

type ChangeBalanceOut Status

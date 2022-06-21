package dto

import "github.com/gookit/validate"

type ChangeBalanceIN struct {
	UserID      int64  `json:"user_id"     validate:"required|gt:0"`
	Amount      int64  `json:"change"      validate:"required|ne:1"`
	Description string `json:"description" validate:"required|min_len:10"`
}

func (c ChangeBalanceIN) Validate() error {
	if v := validate.Struct(c); !v.Validate() {
		return v.Errors.OneError()
	}

	return nil
}

type ChangeBalanceOUT Status

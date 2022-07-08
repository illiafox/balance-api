package dto

import "github.com/gookit/validate"

type TransferBalanceIN struct {
	ToID        int64  `json:"to_id"       validate:"required|gt:0"`
	FromID      int64  `json:"from_id"     validate:"required|gt:0"`
	Amount      int64  `json:"amount"      validate:"required|gt:0"`
	Description string `json:"description" validate:"required"` // |min_len:10
}

func (t TransferBalanceIN) Validate() error {
	if v := validate.Struct(t); !v.Validate() {
		return v.Errors.OneError()
	}

	return nil
}

type TransferBalanceOUT Status

package dto

import (
	"balance-service/app/internal/controller/http/httputils"
	"github.com/gookit/validate"
)

type TransferBalanceIN struct {
	ToID        uint64 `json:"to_id"       validate:"required|gt:0"`
	FromID      uint64 `json:"from_id"     validate:"required|gt:0"`
	Amount      uint64 `json:"amount"      validate:"required|gt:0"`
	Description string `json:"description" validate:"required"` // |min_len:10
}

func (t TransferBalanceIN) Validate() error {
	if v := validate.Struct(t); !v.Validate() {
		return v.Errors.OneError()
	}

	return nil
}

type TransferBalanceOUT httputils.Status

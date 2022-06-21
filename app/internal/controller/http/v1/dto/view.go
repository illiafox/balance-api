package dto

import (
	"balance-service/app/internal/domain/entity"
	"github.com/gookit/validate"
)

type ViewTransactionsIN struct {
	UserID int64  `json:"user_id" validate:"required|gt:0"`
	Sort   string `json:"sort"    validate:"required"`
	Limit  int64  `json:"limit"   validate:"required|gt:0|lte:100"`
	Offset int64  `json:"offset"  validate:"gte:0"`
}

func (c ViewTransactionsIN) Validate() error {
	if v := validate.Struct(c); !v.Validate() {
		return v.Errors.OneError()
	}

	return nil
}

type ViewTransactionsOUT struct {
	Status
	Transactions []entity.Transaction `json:"transactions"`
}

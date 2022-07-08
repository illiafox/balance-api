package dto

import (
	"fmt"
	"net/url"
	"strconv"

	"balance-service/app/internal/domain/entity"
)

type ViewTransactionsIN struct {
	UserID int64  `json:"user_id" validate:"required|gt:0"`
	Sort   string `json:"sort"    validate:"required"`
	Limit  int64  `json:"limit"   validate:"required|gt:0|lte:100"`
	Offset int64  `json:"offset"  validate:"gte:0"`
}

func (v *ViewTransactionsIN) ParseAndValidate(query url.Values) error {
	var err error

	// sort
	v.Sort = query.Get("sort")

	// offset
	if offset := query.Get("offset"); offset != "" {
		if v.Offset, err = strconv.ParseInt(offset, 10, 64); err != nil {
			return fmt.Errorf("parse offset: %w", err)
		}

		if v.Offset < 0 {
			return fmt.Errorf("wrong offset value: %d", v.Offset)
		}
	}

	// limit
	if limit := query.Get("limit"); limit != "" {
		if v.Limit, err = strconv.ParseInt(limit, 10, 64); err != nil {
			return fmt.Errorf("parse offset: %w", err)
		}

		if v.Limit < 0 || v.Limit > 100 {
			return fmt.Errorf("wrong limit value: %d", v.Offset)
		}
	} else {
		v.Limit = 100
	}

	return nil
}

type ViewTransactionsOUT struct {
	Status
	Transactions []entity.Transaction `json:"transactions"`
}

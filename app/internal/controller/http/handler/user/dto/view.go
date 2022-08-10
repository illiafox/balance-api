package dto

import (
	"fmt"
	"net/http"
	"strconv"

	"balance-service/app/internal/controller/http/httputils"
	"balance-service/app/internal/domain/entity"
	"github.com/gin-gonic/gin/binding"
	"github.com/gookit/validate"
)

type ViewTransactionsIN struct {
	Sort   string `query:"sort"`
	Limit  int64  `query:"limit"             validate:"required|gt:0|lte:100"`
	Offset int64  `query:"offset"            validate:"gte:0"`
	UserID int64  `validate:"required|gte:0"`
}

func NewViewTransactionsIN(id string, r *http.Request) (ViewTransactionsIN, error) {
	view := ViewTransactionsIN{
		Limit: 100,
	}

	var err error

	// UserID
	if view.UserID, err = strconv.ParseInt(id, 10, 64); err != nil {
		return view, fmt.Errorf("parse id: %w", err)
	}
	if view.UserID <= 0 {
		return view, fmt.Errorf("id: expected > 0, got %d", view.UserID)
	}

	// Bind
	err = binding.Query.Bind(r, &view)
	if err != nil {
		return view, fmt.Errorf("bind: %w", err)
	}

	// Validate
	if v := validate.Struct(view); !v.Validate() {
		return view, v.Errors.OneError()
	}

	return view, nil
}

// //

type ViewTransactionsOUT struct {
	httputils.Status
	Transactions []entity.Transaction `json:"transactions"`
}

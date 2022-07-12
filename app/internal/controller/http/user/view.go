package user

import (
	"net/http"

	"balance-service/app/internal/controller/http/httputils"
	"balance-service/app/internal/controller/http/middleware"
	"balance-service/app/internal/controller/http/user/dto"
	"balance-service/app/pkg/errors"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

import _ "balance-service/app/internal/domain/entity" // for swagger type recognition

// ViewTransactions
// @Summary      View user transactions
// @Description  View transactions with sorting and pagination
// @Tags         transactions
// @Produce      json
// @Param        id     path    int  	true  "user id"			minimum(1)
// @Param        limit  query   int  	false "output limit"	minimum(0) maximum(100) default(100)
// @Param        offset query   int  	false "output offset" 	minimum(0) default(0)
// @Param        sort	query   string	false  "sort type"  	Enums(DATE_DESC, DATE_ASC, SUM_DESC, SUM_ASC)
// @Success      200  {object}  dto.ViewTransactionsOUT "Transactions data"
// @Failure      422  {object}  httputils.Error
// @Failure      500  {object}  httputils.Error
// @Router       /user/{id}/transactions [get]
func (h *handler) ViewTransactions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context() // context with logger

	// // parse input

	view, err := dto.NewViewTransactionsIN(
		ps.ByName("id"), // path
		r,               // request (url query)
	)

	if err != nil {
		_ = httputils.NewError(w, http.StatusBadRequest, err)
		return
	}

	// // call services

	// get transactions
	t, err := h.balanceService.GetTransactions(ctx, view.UserID, view.Limit, view.Offset, view.Sort)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			middleware.GetLogger(ctx).Error("get transactions",
				zap.Error(err),
				zap.Uint64("user_id", view.UserID),
			)
			_ = httputils.NewError(w, http.StatusInternalServerError, internal)
		} else {
			_ = httputils.NewError(w, http.StatusUnprocessableEntity, err)
		}

		return
	}

	out := dto.ViewTransactionsOUT{
		Status:       httputils.Status{Ok: true},
		Transactions: t,
	}

	// // encode response
	err = httputils.NewResponse(w, out)

	if err != nil {
		middleware.GetLogger(ctx).Error("encode response",
			zap.Error(err), zap.Any("response", out),
		)
		_ = httputils.NewError(w, http.StatusInternalServerError, err)
	}
}

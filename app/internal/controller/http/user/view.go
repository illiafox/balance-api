package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"balance-service/app/internal/controller/http/httputils"
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
// @Success      200  {object}  dto.ViewTransactionsOUT{transactions=[]entity.Transaction{from_id=integer}} "Transactions data"
// @Failure      422  {object}  httputils.Error
// @Failure      500  {object}  httputils.Error
// @Router       /user/{id}/transactions [get]
func (h *handler) ViewTransactions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var trs dto.ViewTransactionsIN

	{
		id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
		if err != nil {
			_ = httputils.NewError(w, http.StatusBadRequest, fmt.Errorf("invalid request query: parse id: %w", err))
			return
		}
		trs.UserID = id
	}

	// validate
	if err := trs.ParseAndValidate(r.URL.Query()); err != nil {
		_ = httputils.NewError(w, http.StatusBadRequest, err)
		return
	}

	// // call service
	ctx, cancel := context.WithTimeout(r.Context(), h.timeout)
	defer cancel()

	t, err := h.balanceService.GetTransactions(ctx, trs.UserID, trs.Limit, trs.Offset, trs.Sort)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			h.logger.Error("get transactions", zap.Error(err), zap.Int64("user_id", trs.UserID))
			_ = httputils.NewError(w, http.StatusInternalServerError, internal)

			return
		} else {
			_ = httputils.NewError(w, http.StatusUnprocessableEntity, err)
		}

		return
	}

	// encode response
	err = httputils.NewResponse(w, dto.ViewTransactionsOUT{
		Status:       dto.Status{Ok: true},
		Transactions: t,
	})

	if err != nil {
		h.logger.Error("/view: encode response", zap.Error(err))
		_ = httputils.NewError(w, http.StatusInternalServerError, err)
	}
}

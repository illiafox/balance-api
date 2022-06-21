package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"balance-service/app/internal/controller/http/v1/dto"
	"balance-service/app/pkg/errors"
	"go.uber.org/zap"
)

// ViewTransactions
// @Summary      View user transactions
// @Description  View transactions with sorting and pagination
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        input body 	dto.ViewTransactionsIN false "User id, limit, sorting and pagination"
// @Success      200  {object}  dto.ViewTransactionsOUT
// @Failure      406  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /view [get]
func (h *handler) ViewTransactions(w http.ResponseWriter, r *http.Request) {
	var trs dto.ViewTransactionsIN
	// decode body
	defer r.Body.Close() // ignore error
	if err := json.NewDecoder(r.Body).Decode(&trs); err != nil {
		dto.JSONError(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))

		return
	}
	// validate
	if err := trs.Validate(); err != nil {
		dto.JSONError(w, http.StatusBadRequest, err)

		return
	}
	//
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()
	// call service
	t, err := h.balanceService.GetTransactions(ctx, trs.UserID, trs.Limit, trs.Offset, trs.Sort)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			h.logger.Error("/view: get transactions", zap.Error(err), zap.Int64("user_id", trs.UserID))
			dto.JSONError(w, http.StatusInternalServerError, internal)

			return
		} else {
			dto.JSONError(w, http.StatusNotAcceptable, err)
		}

		return
	}
	// encode response
	err = dto.JSONResponse(w, dto.ViewTransactionsOUT{
		Status:       dto.Status{Ok: true},
		Transactions: t,
	})
	//
	if err != nil {
		h.logger.Error("/view: encode response", zap.Error(err))
	}
}

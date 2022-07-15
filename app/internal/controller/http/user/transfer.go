package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"balance-service/app/internal/controller/http/httputils"
	"balance-service/app/internal/controller/http/middleware"
	"balance-service/app/internal/controller/http/user/dto"
	"balance-service/app/pkg/errors"
	"go.uber.org/zap"
)

// TransferBalance
// @Summary      Transfer money between users
// @Description  Transfer money from one balance to another
// @Tags         balance
// @Accept       json
// @Produce      json
// @Param        input body 	dto.TransferBalanceIN  true "To and From ID, Amount and Description"
// @Success      200  {object}  dto.TransferBalanceOUT
// @Failure      422  {object}  httputils.Error
// @Failure      500  {object}  httputils.Error
// @Router       /user/transfer [post]
func (h *handler) TransferBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // context with logger

	// // parse input
	var transfer dto.TransferBalanceIN

	// decode json body
	defer r.Body.Close() // ignore error
	if err := json.NewDecoder(r.Body).Decode(&transfer); err != nil {
		_ = httputils.NewError(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))

		return
	}

	// validate struct
	if err := transfer.Validate(); err != nil {
		_ = httputils.NewError(w, http.StatusBadRequest, err)

		return
	}

	// // call services

	// transfer
	err := h.balanceService.Transfer(ctx, transfer.FromID, transfer.ToID, transfer.Amount, transfer.Description)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			middleware.GetLogger(ctx).Error("transfer balance",
				zap.Error(err), zap.Int64("amount", transfer.Amount),
				zap.Int64("from_id", transfer.FromID),
				zap.Int64("to_id", transfer.ToID),
			)
			_ = httputils.NewError(w, http.StatusInternalServerError, internal)
		} else {
			_ = httputils.NewError(w, http.StatusUnprocessableEntity, err)
		}

		return
	}

	out := dto.TransferBalanceOUT{Ok: true}

	// // encode response
	err = httputils.NewResponse(w, out)

	if err != nil {
		middleware.GetLogger(ctx).Error("encode response",
			zap.Error(err), zap.Any("response", out),
		)
		_ = httputils.NewError(w, http.StatusInternalServerError, err)
	}
}

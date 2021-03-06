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

// ChangeBalance
// @Summary      Change user balance
// @Description  Change balance by user ID
// @Tags         balance
// @Accept       json
// @Produce      json
// @Param        input   body 	dto.ChangeBalanceIN		false 	"User ID, Change amount and Description"
// @Success      200  {object}  dto.ChangeBalanceOUT
// @Failure      400  {object}  httputils.Error
// @Failure      422  {object}  httputils.Error
// @Failure      500  {object}  httputils.Error
// @Router       /user/change [patch]
func (h *handler) ChangeBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // context with logger

	// // parse input
	var change dto.ChangeBalanceIN

	// decode body
	defer r.Body.Close() // ignore error

	if err := json.NewDecoder(r.Body).Decode(&change); err != nil {
		_ = httputils.NewError(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))
		return
	}

	// validate struct
	if err := change.Validate(); err != nil {
		_ = httputils.NewError(w, http.StatusBadRequest, err)
		return
	}

	// // call services

	// change balance
	err := h.balanceService.Change(ctx, change.UserID, change.Amount, change.Description)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			middleware.GetLogger(ctx).Error("change balance",
				zap.Error(err),
				zap.Int64("user_id", change.UserID),
				zap.Int64("amount", change.Amount),
			)
			_ = httputils.NewError(w, http.StatusInternalServerError, internal)
		} else {
			_ = httputils.NewError(w, http.StatusUnprocessableEntity, err)
		}

		return
	}

	out := dto.ChangeBalanceOUT{Ok: true}

	// // encode response
	err = httputils.NewResponse(w, out)

	if err != nil {
		middleware.GetLogger(ctx).Error("encode response",
			zap.Error(err), zap.Any("response", out),
		)
		_ = httputils.NewError(w, http.StatusInternalServerError, err)
	}
}

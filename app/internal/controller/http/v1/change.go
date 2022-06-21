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

// ChangeBalance
// @Summary      Change user balance
// @Description  Change balance by user ID
// @Tags         balance
// @Accept       json
// @Produce      json
// @Param        input body 	dto.ChangeBalanceIN false "User id, change amount and description"
// @Success      200  {object}  dto.ChangeBalanceOUT
// @Failure      400  {object}  dto.Error
// @Failure      406  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /change [post]
func (h *handler) ChangeBalance(w http.ResponseWriter, r *http.Request) {
	var change dto.ChangeBalanceIN
	// decode body
	defer r.Body.Close() // ignore error
	if err := json.NewDecoder(r.Body).Decode(&change); err != nil {
		dto.JSONError(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))

		return
	}
	// validate struct
	if err := change.Validate(); err != nil {
		dto.JSONError(w, http.StatusBadRequest, err)

		return
	}
	//
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()
	// call service
	err := h.balanceService.Change(ctx, change.UserID, change.Amount, change.Description)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			h.logger.Error("/change: change balance", zap.Error(err), zap.Int64("user_id", change.UserID), zap.Int64("amount", change.Amount))
			dto.JSONError(w, http.StatusInternalServerError, internal)
		} else {
			dto.JSONError(w, http.StatusNotAcceptable, err)
		}

		return
	}
	// encode response
	err = dto.JSONResponse(w, dto.ChangeBalanceOUT{Ok: true})
	//
	if err != nil {
		h.logger.Error("/change: encode response", zap.Error(err))
	}
}

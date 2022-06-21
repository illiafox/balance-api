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

// dungeon

// TransferBalance
// @Summary      Transfer money between users
// @Description  Transfer money from one balance to another
// @Tags         balance
// @Accept       json
// @Produce      json
// @Param        input body 	dto.TransferBalanceIN false "to, from, amount and description"
// @Success      200  {object}  dto.TransferBalanceOUT
// @Failure      406  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /transfer [put]
func (h *handler) TransferBalance(w http.ResponseWriter, r *http.Request) {
	var transfer dto.TransferBalanceIN
	// decode body
	defer r.Body.Close() // ignore error
	if err := json.NewDecoder(r.Body).Decode(&transfer); err != nil {
		dto.JSONError(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))

		return
	}
	// validate struct
	if err := transfer.Validate(); err != nil {
		dto.JSONError(w, http.StatusBadRequest, err)

		return
	}
	//
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()
	// call service
	err := h.balanceService.Transfer(ctx, transfer.FromID, transfer.ToID, transfer.Amount, transfer.Description)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			h.logger.Error("/transfer: transfer balance",
				zap.Error(err), zap.Int64("amount", transfer.Amount),
				zap.Int64("from_id", transfer.FromID),
				zap.Int64("to_id", transfer.ToID),
			)
			dto.JSONError(w, http.StatusInternalServerError, internal)
		} else {
			dto.JSONError(w, http.StatusNotAcceptable, err)
		}

		return
	}
	// encode response
	err = dto.JSONResponse(w, dto.TransferBalanceOUT{Ok: true})
	//
	if err != nil {
		h.logger.Error("/transfer: encode response", zap.Error(err))
	}
}

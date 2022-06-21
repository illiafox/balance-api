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

// GetBalance
// @Summary      Get user balance
// @Description  Get balance by user ID
// @Tags         balance
// @Accept       json
// @Produce      json
// @Param        input body 	dto.GetBalanceIN false "User id and Currency"
// @Success      200  {object}  dto.GetBalanceOUT{balance=integer} "Balance data"
// @Failure      400  {object}  dto.Error
// @Failure      406  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /get [get]
func (h *handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	var get dto.GetBalanceIN
	// decode body
	defer r.Body.Close() // ignore error
	if err := json.NewDecoder(r.Body).Decode(&get); err != nil {
		dto.JSONError(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))

		return
	}
	// validate
	if err := get.Validate(); err != nil {
		dto.JSONError(w, http.StatusBadRequest, err)

		return
	}
	//
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()
	// call service
	balance, err := h.balanceService.Get(ctx, get.UserID, get.Base)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			h.logger.Error("/get: get balance", zap.Error(err), zap.Int64("user_id", get.UserID), zap.String("base", get.Base))
			dto.JSONError(w, http.StatusInternalServerError, internal)

			return
		} else {
			dto.JSONError(w, http.StatusNotAcceptable, err)
		}

		return
	}
	// encode response
	err = dto.JSONResponse(w, dto.GetBalanceOUT{
		Status:  dto.Status{Ok: true},
		Base:    get.Base,
		Balance: []byte(balance),
	})
	//
	if err != nil {
		h.logger.Error("/get: encode response", zap.Error(err), zap.String("balance", balance))
	}
}

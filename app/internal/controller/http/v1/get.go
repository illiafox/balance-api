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

func (h *handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	logger := h.logger.Named("get")
	//
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()
	//
	var get dto.GetBalanceIn
	// decode body
	defer r.Body.Close() // ignore error
	if err := json.NewDecoder(r.Body).Decode(&get); err != nil {
		dto.JSONError(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))

		return
	}
	// check user id
	if get.UserID <= 0 {
		dto.JSONError(w, http.StatusBadRequest, fmt.Errorf("invalid user id: got %d, expected > 0", get.UserID))

		return
	}
	// call service
	balance, err := h.balanceService.Get(ctx, get.UserID, get.Base)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			logger.Error("get balance", zap.Error(err), zap.Int64("user_id", get.UserID), zap.String("base", get.Base))
			dto.JSONError(w, http.StatusInternalServerError, internal)

			return
		}

		dto.JSONError(w, http.StatusNotAcceptable, err)
	}
	// encode response
	err = dto.JSONResponse(w, dto.GetBalanceOut{
		Status:  dto.Status{Ok: true},
		Balance: []byte(balance),
		Base:    get.Base,
	})
	//
	if err != nil {
		logger.Error("encode response", zap.Error(err), zap.String("balance", balance), zap.String("base", get.Base))
		dto.JSONError(w, http.StatusInternalServerError, err)
	}
}

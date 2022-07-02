package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"balance-service/app/internal/controller/http/v1/user/dto"
	"balance-service/app/pkg/errors"
	"go.uber.org/zap"
)

// BlockBalance
// @Summary      Block user balance
// @Description  Block balance by user ID
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        input body 	dto.BlockIN		true 	"User ID and Reason"
// @Success      200  {object}  dto.BlockOUT
// @Failure      400  {object}  dto.Error
// @Failure      422  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /user/block [post]
func (h *handler) BlockBalance(w http.ResponseWriter, r *http.Request) {

	var block dto.BlockIN

	// decode body
	defer r.Body.Close() // ignore error
	if err := json.NewDecoder(r.Body).Decode(&block); err != nil {
		dto.JSONError(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))

		return
	}

	// validate struct
	if err := block.Validate(); err != nil {
		dto.JSONError(w, http.StatusBadRequest, err)

		return
	}

	// // call service
	ctx, cancel := context.WithTimeout(r.Context(), h.timeout)
	defer cancel()

	err := h.balanceService.BlockBalance(ctx, block.UserID, block.Reason)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			h.logger.Error("/block: block balance", zap.Error(err), zap.Int64("user_id", block.UserID))
			dto.JSONError(w, http.StatusInternalServerError, internal)
		} else {
			dto.JSONError(w, http.StatusUnprocessableEntity, err)
		}

		return
	}

	if err = dto.JSONResponse(w, dto.BlockOUT{Ok: true}); err != nil {
		h.logger.Error("/block: encode response", zap.Error(err))
		dto.JSONError(w, http.StatusInternalServerError, err)
	}
}

// UnblockBalance
// @Summary      Unblock user balance
// @Description  Unblock balance by user ID
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        input body 	dto.UnblockIN		true 	"User ID"
// @Success      200  {object}  dto.UnblockOUT
// @Failure      400  {object}  dto.Error
// @Failure      422  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /user/unblock [post]
func (h *handler) UnblockBalance(w http.ResponseWriter, r *http.Request) {

	var unblock dto.UnblockIN

	// decode body
	defer r.Body.Close() // ignore error
	if err := json.NewDecoder(r.Body).Decode(&unblock); err != nil {
		dto.JSONError(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))

		return
	}

	// validate struct
	if err := unblock.Validate(); err != nil {
		dto.JSONError(w, http.StatusBadRequest, err)

		return
	}

	// // call service
	ctx, cancel := context.WithTimeout(r.Context(), h.timeout)
	defer cancel()

	err := h.balanceService.UnblockBalance(ctx, unblock.UserID)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			h.logger.Error("/unblock: unblock balance", zap.Error(err), zap.Int64("user_id", unblock.UserID))
			dto.JSONError(w, http.StatusInternalServerError, internal)
		} else {
			dto.JSONError(w, http.StatusUnprocessableEntity, err)
		}

		return
	}

	if err = dto.JSONResponse(w, dto.UnblockOUT{Ok: true}); err != nil {
		h.logger.Error("/unblock: encode response", zap.Error(err))
		dto.JSONError(w, http.StatusInternalServerError, err)
	}
}

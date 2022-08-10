package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"balance-service/app/internal/controller/http/handler/admin/dto"
	"balance-service/app/internal/controller/http/httputils"
	"balance-service/app/internal/controller/http/middleware"
	"balance-service/app/pkg/errors"
	"balance-service/app/pkg/logger"
)

// BlockBalance
// @Summary      Block user balance
// @Description  Block balance by user ID
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        input body 	dto.BlockIN		true 	"User ID and Reason"
// @Success      200  {object}  dto.BlockOUT
// @Failure      400  {object}  httputils.Error
// @Failure      422  {object}  httputils.Error
// @Failure      500  {object}  httputils.Error
// @Router       /admin/block [post]
func (h *handler) BlockBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // context with logger

	// // parse input
	var block dto.BlockIN

	// decode json
	defer r.Body.Close() // ignore error
	if err := json.NewDecoder(r.Body).Decode(&block); err != nil {
		_ = httputils.NewError(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))

		return
	}

	// validate struct
	if err := block.Validate(); err != nil {
		_ = httputils.NewError(w, http.StatusBadRequest, err)
		return
	}

	// // call services

	// block balance
	err := h.balanceService.BlockBalance(ctx, block.UserID, block.Reason)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			middleware.GetLogger(ctx).Error("block balance",
				logger.Error(err), logger.Int64("user_id", block.UserID),
			)
			_ = httputils.NewError(w, http.StatusInternalServerError, internal)
		} else {
			_ = httputils.NewError(w, http.StatusUnprocessableEntity, err)
		}

		return
	}

	out := dto.BlockOUT{Ok: true}

	// // encode response
	err = httputils.NewResponse(w, out)

	if err != nil {
		middleware.GetLogger(ctx).Error("encode response",
			logger.Error(err), logger.Any("response", out),
		)
		_ = httputils.NewError(w, http.StatusInternalServerError, err)
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
// @Failure      400  {object}  httputils.Error
// @Failure      422  {object}  httputils.Error
// @Failure      500  {object}  httputils.Error
// @Router       /admin/unblock [post]
func (h *handler) UnblockBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // context with logger

	// // parse input
	var unblock dto.UnblockIN

	// decode json
	defer r.Body.Close() // ignore error
	if err := json.NewDecoder(r.Body).Decode(&unblock); err != nil {
		_ = httputils.NewError(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))

		return
	}

	// validate struct
	if err := unblock.Validate(); err != nil {
		_ = httputils.NewError(w, http.StatusBadRequest, err)
		return
	}

	// // call services

	// unblock
	err := h.balanceService.UnblockBalance(ctx, unblock.UserID)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			middleware.GetLogger(ctx).Error("unblock balance",
				logger.Error(err), logger.Int64("user_id", unblock.UserID),
			)
			_ = httputils.NewError(w, http.StatusInternalServerError, internal)
		} else {
			_ = httputils.NewError(w, http.StatusUnprocessableEntity, err)
		}

		return
	}

	out := dto.UnblockOUT{Ok: true}

	// // encode response
	err = httputils.NewResponse(w, out)

	if err != nil {
		middleware.GetLogger(ctx).Error("encode response",
			logger.Error(err), logger.Any("response", out),
		)
		_ = httputils.NewError(w, http.StatusInternalServerError, err)
	}
}

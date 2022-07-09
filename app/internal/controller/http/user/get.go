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

// GetBalance
// @Summary      Get user balance
// @Description  Get balance by User ID
// @Tags         balance
// @Produce      json
// @Param        id     path    int  	true  	"user id" 				minimum(1)
// @Param        base	query   string	false  	"currency abbreviation" minlength(3)
// @Success      200  {object}  dto.GetBalanceOUT{balance=integer} "Balance data"
// @Failure      400  {object}  httputils.Error
// @Failure      406  {object}  httputils.Error
// @Failure      500  {object}  httputils.Error
// @Router       /user/{id} [get]
func (h *handler) GetBalance(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// read query values
	var get = dto.GetBalanceIN{
		Base: r.URL.Query().Get("base"),
	}

	// path parameter
	var err error
	if get.UserID, err = strconv.ParseInt(ps.ByName("id"), 10, 64); err != nil {
		_ = httputils.NewError(w, http.StatusBadRequest, fmt.Errorf("invalid request query: parse id: %w", err))
		return
	}

	// validate
	if err := get.Validate(); err != nil {
		_ = httputils.NewError(w, http.StatusBadRequest, err)
		return
	}

	// // call service
	ctx, cancel := context.WithTimeout(r.Context(), h.timeout)
	defer cancel()

	balance, err := h.balanceService.Get(ctx, get.UserID, get.Base)
	if err != nil {

		if internal, ok := errors.ToInternal(err); ok {
			h.logger.Error("get balance", zap.Error(err), zap.Int64("user_id", get.UserID), zap.String("base", get.Base))
			_ = httputils.NewError(w, http.StatusInternalServerError, internal)
		} else {
			_ = httputils.NewError(w, http.StatusNotAcceptable, err)
		}

		return
	}

	// // encode response
	err = httputils.NewResponse(w, dto.GetBalanceOUT{
		Status:  dto.Status{Ok: true},
		Base:    get.Base,
		Balance: dto.Balance(balance),
	})
	if err != nil {
		h.logger.Error("encode response", zap.Error(err), zap.String("balance", balance))
		_ = httputils.NewError(w, http.StatusInternalServerError, err)
	}
}

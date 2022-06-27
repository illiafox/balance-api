package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"balance-service/app/internal/controller/http/v1/dto"
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
// @Failure      400  {object}  dto.Error
// @Failure      406  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /user/{id} [get]
func (h *handler) GetBalance(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// read query values
	var get = dto.GetBalanceIN{
		Base: r.URL.Query().Get("base"),
	}

	// path parameter
	var err error
	if get.UserID, err = strconv.ParseInt(ps.ByName("id"), 10, 64); err != nil {
		dto.JSONError(w, http.StatusBadRequest, fmt.Errorf("invalid request query: parse id: %w", err))

		return
	}

	// validate
	if err := get.Validate(); err != nil {
		dto.JSONError(w, http.StatusBadRequest, err)

		return
	}

	// // call service
	ctx, cancel := context.WithTimeout(r.Context(), h.timeout)
	defer cancel()

	balance, err := h.balanceService.Get(ctx, get.UserID, get.Base)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			h.logger.Error("get balance", zap.Error(err), zap.Int64("user_id", get.UserID), zap.String("base", get.Base))
			dto.JSONError(w, http.StatusInternalServerError, internal)

			return
		} else {
			dto.JSONError(w, http.StatusNotAcceptable, err)
		}

		return
	}

	// // encode response
	err = dto.JSONResponse(w, dto.GetBalanceOUT{
		Status:  dto.Status{Ok: true},
		Base:    get.Base,
		Balance: []byte(balance),
	})
	if err != nil {
		h.logger.Error("encode response", zap.Error(err), zap.String("balance", balance))
		dto.JSONError(w, http.StatusInternalServerError, err)
	}
}

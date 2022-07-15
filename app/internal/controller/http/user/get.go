package user

import (
	"net/http"

	"balance-service/app/internal/controller/http/httputils"
	"balance-service/app/internal/controller/http/middleware"
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
	ctx := r.Context() // context with logger

	// // parse input
	get, err := dto.NewGetBalanceIN(
		ps.ByName("id"),           // path
		r.URL.Query().Get("base"), // query
	)

	if err != nil {
		_ = httputils.NewError(w, http.StatusBadRequest, err)
		return
	}

	// // call services

	// get balance
	balance, err := h.balanceService.Get(ctx, get.UserID, get.Base)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			_ = httputils.NewError(w, http.StatusInternalServerError, internal)
		} else {
			_ = httputils.NewError(w, http.StatusNotAcceptable, err)
		}

		return
	}

	out := dto.GetBalanceOUT{
		Status:  httputils.Status{Ok: true},
		Base:    get.Base,
		Balance: dto.Balance(balance),
	}

	// // encode response
	err = httputils.NewResponse(w, out)

	if err != nil {
		middleware.GetLogger(ctx).Error("encode response",
			zap.Error(err),
			zap.Any("response", out),
		)
		_ = httputils.NewError(w, http.StatusInternalServerError, err)
	}
}

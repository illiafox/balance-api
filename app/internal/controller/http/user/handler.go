package user

import (
	"net/http"

	"balance-service/app/internal/adapters/api"
	"balance-service/app/internal/adapters/api/balance"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	balanceService balance.Service
}

func New(balanceService balance.Service) api.Handler {
	return &handler{
		balanceService: balanceService,
	}
}

func (h *handler) Handler() http.Handler {
	router := httprouter.New()

	router.GET("/:id", h.GetBalance)
	router.GET("/:id/transactions", h.ViewTransactions)
	//
	router.PATCH("/change", wrap(h.ChangeBalance))
	router.POST("/transfer", wrap(h.TransferBalance))

	return router
}

func wrap(f http.HandlerFunc) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		f(writer, request)
	}
}

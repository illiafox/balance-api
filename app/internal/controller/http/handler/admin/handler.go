package admin

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

	router.POST("/block", wrap(h.BlockBalance))
	router.POST("/unblock", wrap(h.UnblockBalance))

	return router
}

func wrap(f http.HandlerFunc) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		f(writer, request)
	}
}

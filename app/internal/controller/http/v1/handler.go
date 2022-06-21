package v1

import (
	"net/http"
	"time"

	"balance-service/app/internal/adapters/api"
	"balance-service/app/internal/adapters/api/balance"
	"balance-service/app/pkg/logger"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	balanceService balance.Service
	//
	logger logger.Logger
	//
	timeout time.Duration
}

func NewHandler(logger logger.Logger, balanceService balance.Service) api.Handler {
	return &handler{
		balanceService: balanceService,
		//
		logger: logger,
		//
		timeout: time.Second,
	}
}

func (h *handler) Register() http.Handler {
	router := httprouter.New()

	router.GET("/get", wrap(h.GetBalance))
	router.POST("/change", wrap(h.ChangeBalance))
	router.PUT("/transfer", wrap(h.TransferBalance))
	router.GET("/view", wrap(h.ViewTransactions))

	return router
}

func wrap(f http.HandlerFunc) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		f(writer, request)
	}
}

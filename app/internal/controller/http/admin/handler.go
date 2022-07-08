package admin

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

func New(logger logger.Logger, balanceService balance.Service) api.Handler {
	return &handler{
		balanceService: balanceService,
		//
		logger: logger,
		//
		timeout: time.Second * 5,
	}
}

func (h *handler) Handler() http.Handler {
	user := httprouter.New()

	user.POST("/block", wrap(h.BlockBalance))
	user.POST("/unblock", wrap(h.UnblockBalance))

	return user
}

func wrap(f http.HandlerFunc) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		f(writer, request)
	}
}
package v1

import (
	"net/http"
	"time"

	"balance-service/app/internal/adapters/api"
	"balance-service/app/internal/adapters/api/balance"
	"go.uber.org/zap"
)

type handler struct {
	balanceService balance.Service
	//
	logger *zap.Logger
	//
	timeout time.Duration
}

func NewHandler(logger *zap.Logger, balanceService balance.Service) api.Handler {
	return &handler{
		balanceService: balanceService,
		//
		logger: logger,
		//
		timeout: time.Second,
	}
}

func (h *handler) Register(router *http.ServeMux) {

}

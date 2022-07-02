package v1

import (
	"net/http"

	"balance-service/app/internal/adapters/api/balance"
	"balance-service/app/internal/controller/http/v1/user"
	"balance-service/app/pkg/logger"
)

func New(logger logger.Logger, balanceService balance.Service) http.Handler {
	router := http.NewServeMux()

	{
		u := user.New(logger, balanceService).Handler()
		router.Handle("/user/", http.StripPrefix("/user", u))
	}

	return router
}

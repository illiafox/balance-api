package handler

import (
	"net/http"
	"time"

	"balance-service/app/internal/adapters/api/balance"
	"balance-service/app/internal/controller/http/handler/admin"
	"balance-service/app/internal/controller/http/handler/user"
	"balance-service/app/internal/controller/http/middleware"
	"balance-service/app/pkg/logger"
)

func New(logger logger.Logger, timeout time.Duration, balanceService balance.Service) http.Handler {
	router := http.NewServeMux()
	//
	{
		// user
		usr := middleware.New(logger.Named("user"), timeout).Use(
			user.New(balanceService).Handler(),
		)
		router.Handle("/user/", http.StripPrefix("/user", usr))
	}
	//
	{
		// admin
		adm := middleware.New(logger.Named("admin"), timeout).Use(
			admin.New(balanceService).Handler(),
		)
		router.Handle("/admin/", http.StripPrefix("/admin", adm))
	}
	//
	return router
}

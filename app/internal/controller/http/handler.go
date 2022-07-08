package http

import (
	"net/http"

	"balance-service/app/internal/adapters/api/balance"
	"balance-service/app/internal/controller/http/admin"
	"balance-service/app/internal/controller/http/user"
	"balance-service/app/pkg/logger"
)

func New(logger logger.Logger, balanceService balance.Service) http.Handler {
	router := http.NewServeMux()

	u := user.New(logger.Named("user"), balanceService).Handler()
	router.Handle("/user/", http.StripPrefix("/user", u))
	//
	a := admin.New(logger.Named("admin"), balanceService).Handler()
	router.Handle("/admin/", http.StripPrefix("/admin", a))

	return router
}

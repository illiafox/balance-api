package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ozon-url-shortener/app/internal/adapters/links/api"
	"ozon-url-shortener/app/internal/domain/links"
)

func (app *App) Handler() http.Handler {
	h := api.NewHandler(
		app.logger.Named("api"),
		links.NewService(app.Storage()),
	)

	if !app.flags.debug {
		gin.SetMode(gin.ReleaseMode)
	}

	g := gin.New()
	g.Use(gin.Recovery())

	if app.flags.debug {
		g.Use(gin.Logger())
	}

	h.Register(g)

	return g
}

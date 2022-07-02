package api

import "net/http"

type Handler interface {
	Handler() http.Handler
}

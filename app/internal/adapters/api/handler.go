package api

import "net/http"

type Handler interface {
	Register() http.Handler
}

package api

import (
	"net/http"
)

type Handler interface {
	Register(router *http.ServeMux)
}

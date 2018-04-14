package api1

import (
	"log"
	"net/http"

	"goji.io"

	"github.com/canhlinh/go-api/src/services"
)

type api1 struct {
	Srv *services.Srv
	Mux *goji.Mux
}

func (api *api1) DefaultHandler(f services.HandlerFunc) http.Handler {
	return NewHandler(f, api.Srv)
}

func (api *api1) UserLoginRequiredHandler(f services.HandlerFunc) http.Handler {
	h := NewHandler(f, api.Srv)
	h.UserRequired = true
	return h
}

type handler struct {
	handleFunc   func(*services.Context) services.RenderFunc
	Srv          *services.Srv
	UserRequired bool
}

func NewHandler(f services.HandlerFunc, srv *services.Srv) *handler {
	return &handler{
		Srv:        srv,
		handleFunc: f,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			http.Error(w, http.StatusText(500), 500)
		}
	}()

	c := services.NewContext(w, r, h.Srv)
	renderFunc := h.handleFunc(c)
	if renderFunc == nil {
		panic("RenderFunc can not be null")
	}

	if err := renderFunc(c.ResponseCode, c.ResponseData); err != nil {
		log.Println(err)
	}
}

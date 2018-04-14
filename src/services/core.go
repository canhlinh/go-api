package services

import (
	"github.com/canhlinh/go-api/src/config"
	"github.com/canhlinh/go-api/src/stores"
	"github.com/fvbock/endless"
	goji "goji.io"
)

type Srv struct {
	Router *goji.Mux
	Store  stores.Store
}

func NewServer(router *goji.Mux, store stores.Store) *Srv {
	return &Srv{
		Router: router,
		Store:  store,
	}
}

func (srv *Srv) Run() {
	endless.ListenAndServe(config.Config().Server.ListenAddress, srv.Router)
}

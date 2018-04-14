package main

import (
	"flag"

	"github.com/canhlin/go-api/config"
	"github.com/canhlinh/go-api/src/api1"
	"github.com/canhlinh/go-api/src/services"
	"github.com/canhlinh/go-api/src/stores"
	"github.com/canhlinh/go-api/src/utils"
	"goji.io"
)

var srv *services.Srv

func Init() {
	var configPath string
	flag.StringVar(&configPath, "conf", "", "path of the config file")
	flag.Parse()

	if configPath == "" {
		configPath = "./conf/config.yaml"
	}

	config.Load("./conf/config.yaml")
	utils.Init("./i18n")

	srv = services.NewServer(goji.NewMux(), stores.NewStore())
	api1.Init(srv)
}

func StartServer() {
	srv.Run()
}

func main() {
	Init()
	StartServer()
}

// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

import (
	"ginana/internal/config"
	"ginana/internal/db"
	"ginana/internal/server"
	"ginana/internal/server/controller/api"
	"ginana/internal/server/router"
	"ginana/internal/service"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitApp() (*App, func(), error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := db.NewDB(configConfig)
	if err != nil {
		return nil, nil, err
	}
	memcache, err := db.NewMC(configConfig)
	if err != nil {
		return nil, nil, err
	}
	v, err := service.NewErrHelper()
	if err != nil {
		return nil, nil, err
	}
	serviceService, err := service.New(configConfig, gormDB, memcache, v)
	if err != nil {
		return nil, nil, err
	}
	cApi := api.New(serviceService)
	application, err := router.InitRouter(cApi, configConfig)
	if err != nil {
		return nil, nil, err
	}
	httpServer, err := server.NewHttpServer(application, configConfig)
	if err != nil {
		return nil, nil, err
	}
	syncedEnforcer, err := db.NewCasbin(serviceService, configConfig)
	if err != nil {
		return nil, nil, err
	}
	app, cleanup, err := NewApp(httpServer, serviceService, syncedEnforcer)
	if err != nil {
		return nil, nil, err
	}
	return app, func() {
		cleanup()
	}, nil
}

// wire.go:

var initProvider = wire.NewSet(config.NewConfig, db.NewDB, db.NewMC)

var svcProvider = wire.NewSet(service.NewErrHelper, service.New, db.NewCasbin)

var cProvider = wire.NewSet(api.New)

var httpProvider = wire.NewSet(router.InitRouter, server.NewHttpServer)

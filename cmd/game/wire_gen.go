// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"game/internal/biz"
	"game/internal/conf"
	"game/internal/data"
	"game/internal/server"
	"game/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	client := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db, client)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	transaction := data.NewTransaction(dataData)
	appUsecase := biz.NewAppUsecase(userRepo, transaction, logger)
	appService := service.NewAppService(appUsecase, logger, auth)
	httpServer := server.NewHTTPServer(confServer, appService, logger)
	app := newApp(logger, httpServer)
	return app, func() {
		cleanup()
	}, nil
}

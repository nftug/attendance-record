// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"attendance-record/client"
	"attendance-record/client/model"
	"attendance-record/domain/service"
	"attendance-record/infrastructure"
	"attendance-record/infrastructure/repository"
	"attendance-record/shared"
	"attendance-record/usecase"
)

// Injectors from wire.go:

func initApp() *shared.App {
	db := infrastructure.NewDB()
	iWorkRepository := repository.NewWorkRepository(db)
	iRestRepository := repository.NewRestRepository(db)
	timeStatusService := service.NewTimeStatusService(iWorkRepository, iRestRepository)
	timeStatusUseCase := usecase.NewTimeStatusUseCase(timeStatusService)
	app := shared.NewAppSingleton(timeStatusUseCase)
	return app
}

func initClient() *client.Client {
	db := infrastructure.NewDB()
	iWorkRepository := repository.NewWorkRepository(db)
	iRestRepository := repository.NewRestRepository(db)
	timeStatusService := service.NewTimeStatusService(iWorkRepository, iRestRepository)
	timeStatusUseCase := usecase.NewTimeStatusUseCase(timeStatusService)
	app := shared.NewAppSingleton(timeStatusUseCase)
	iTimeStatusApi := model.NewLocalApi(app)
	timeStatusReceiver := model.NewTimeStatusReceiverSingleton(iTimeStatusApi)
	clientClient := client.NewClient(timeStatusReceiver)
	return clientClient
}

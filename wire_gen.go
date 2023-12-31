// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"attendance-record/client"
	"attendance-record/client/model"
	"attendance-record/domain/interfaces"
	"attendance-record/domain/service"
	"attendance-record/infrastructure"
	"attendance-record/infrastructure/localpath"
	"attendance-record/infrastructure/repository"
	"attendance-record/shared"
	"attendance-record/usecase"
)

// Injectors from wire.go:

func initApp() *shared.App {
	localPathService := localpath.NewLocalPathService()
	db := infrastructure.NewDBSingleton(localPathService)
	iWorkRepository := repository.NewWorkSqlRepository(db)
	iRestRepository := repository.NewRestSqlRepository(db)
	timeStatusRepositorySet := interfaces.NewTimeStatusRepositorySet(iWorkRepository, iRestRepository)
	timeStatusService := service.NewTimeStatusService(timeStatusRepositorySet)
	iConfigRepository := repository.NewConfigRepository(localPathService)
	timeStatusUseCase := usecase.NewTimeStatusUseCase(timeStatusService, timeStatusRepositorySet, iConfigRepository)
	configUseCase := usecase.NewConfigUseCase(iConfigRepository)
	app := shared.NewAppSingleton(timeStatusUseCase, configUseCase)
	return app
}

func initClient() *client.Client {
	localPathService := localpath.NewLocalPathService()
	db := infrastructure.NewDBSingleton(localPathService)
	iWorkRepository := repository.NewWorkSqlRepository(db)
	iRestRepository := repository.NewRestSqlRepository(db)
	timeStatusRepositorySet := interfaces.NewTimeStatusRepositorySet(iWorkRepository, iRestRepository)
	timeStatusService := service.NewTimeStatusService(timeStatusRepositorySet)
	iConfigRepository := repository.NewConfigRepository(localPathService)
	timeStatusUseCase := usecase.NewTimeStatusUseCase(timeStatusService, timeStatusRepositorySet, iConfigRepository)
	configUseCase := usecase.NewConfigUseCase(iConfigRepository)
	app := shared.NewAppSingleton(timeStatusUseCase, configUseCase)
	iTimeStatusApi := model.NewTimeStatusLocalApi(app)
	iConfigApi := model.NewConfigLocalApi(app)
	appContainer := model.NewAppContainer(iTimeStatusApi, iConfigApi, localPathService)
	clientClient := client.NewClient(appContainer)
	return clientClient
}

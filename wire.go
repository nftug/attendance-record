//go:build wireinject
// +build wireinject

package main

import (
	"attendance-record/client"
	"attendance-record/domain"
	"attendance-record/infrastructure"
	"attendance-record/shared"
	"attendance-record/usecase"

	"github.com/google/wire"
)

func initApp() *shared.App {
	wire.Build(shared.Set, usecase.Set, domain.Set, infrastructure.Set)
	return nil
}

func initClient(a *shared.App) *client.Client {
	wire.Build(client.Set)
	return nil
}

//go:build wireinject
// +build wireinject

package main

import (
	"attendance-record/client"
	"attendance-record/shared"

	"github.com/google/wire"
)

func initApp() *shared.App {
	wire.Build(shared.Set)
	return nil
}

func initClient() *client.Client {
	wire.Build(client.Set)
	return nil
}

//go:build wireinject
// +build wireinject

package main

import (
	"attendance-record/shared"

	"github.com/google/wire"
)

func initApp() *shared.App {
	wire.Build(shared.Set)
	return nil
}

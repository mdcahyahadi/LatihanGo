//go:build wireinject
// +build wireinject

package uploadphoto

import (
	"latihan/app"

	"github.com/google/wire"
)

var (
	ModuleSets = wire.NewSet(
		NewRepository,
		wire.Bind(new(IRepository), new(*Repository)),
		NewService,
		wire.Bind(new(IService), new(*Service)),
		NewDelivery,
		NewRoutes)
)

func InjectRoutes() (*Routes, error) {
	panic(wire.Build(app.AppModule, ModuleSets))
}

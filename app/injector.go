//go:build wireinject
// +build wireinject

package app

import (
	"latihan/infrastructure/activityLog"
	"latihan/infrastructure/appLog"
	"latihan/infrastructure/configuration"
	"latihan/infrastructure/environment"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var (
	AppModule = wire.NewSet(
		configModuleSets,
		ProvideAppEnvironment,
		ProvideAppEnvConfig,
		ProvideLogger,
		ProvidePostgre,
		ProvideActivityLogger,
	)
)

func InjectAppEnvironment() (environment.AppEnvironment, error) {
	panic(wire.Build(AppModule))
}

func InjectAppConfig() configuration.AppConfig {
	panic(wire.Build(AppModule))
}

func InjectLogger() (*appLog.AppLog, error) {
	panic(wire.Build(AppModule))
}

func InjectPostgre() *gorm.DB {
	panic(wire.Build(AppModule))
}

func InjectActivityLogger() activityLog.IService {
	panic(wire.Build(AppModule))
}

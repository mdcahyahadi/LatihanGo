package app

import (
	"latihan/infrastructure/activityLog"
	"latihan/infrastructure/appLog"
	"latihan/infrastructure/configuration"
	"latihan/infrastructure/database"
	"latihan/infrastructure/environment"

	"gorm.io/gorm"
)

// ProvideAppEnvironment is a function to provide app environment data
func ProvideAppEnvironment() (environment.AppEnvironment, error) {
	return environment.FromOsEnv()
}

// ProvideAppEnvConfig is a function to get AppConfig struct data
func ProvideAppEnvConfig(conf *configuration.AppConfig) configuration.AppConfig {
	return *conf
}

// ProvideLogger is a function to log http request and deployment to a file
func ProvideLogger(env environment.AppEnvironment) *appLog.AppLog {
	return appLog.New(env, appLog.FileTemplate("mygram-svc-%Y_%m_%d"))
}

// ProvidePostgre is function to init postgres connection
func ProvidePostgre(config *configuration.AppConfig) *gorm.DB {
	return database.NewPostgre(config)
}

// ProvideActivityLogger is a function to log activity to db and kafka
func ProvideActivityLogger(conf *configuration.AppConfig) activityLog.IService {
	return activityLog.NewService(conf)
}

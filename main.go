package main

import (
	"latihan/app"
	"latihan/infrastructure/configuration"
	"latihan/infrastructure/environment"
	uploadphoto "latihan/uploadPhoto"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin/v2"
)

func main() {

	// initialize config
	config, err := app.InitializeAppConfig()
	if err != nil {
		panic(err.Error())
	}

	logrus.Info("############################")
	logrus.Info("MyGram Services: ")
	logrus.Info("Powered by : COBACOBA", app.Version)
	logrus.Info("You app running on ", config.App.Environment, " mode")
	logrus.Info("############################")

	if config.App.Environment == string(environment.PROD) {
		gin.SetMode(gin.ReleaseMode)
	}

	// Force log's color
	gin.ForceConsoleColor()
	router := gin.New()
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// APM Integrate
	router.Use(apmgin.Middleware(router))

	// mygram routes
	myGramRoutes, err := uploadphoto.InjectRoutes()
	if err != nil {
		panic(err.Error())
	}
	myGramRoutes.RegisterRoutes(router)

	// r.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(router.Run(":" + config.App.Port))

	configuration.Listen(*config, router)
}

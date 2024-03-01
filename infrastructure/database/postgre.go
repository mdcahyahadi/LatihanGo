package database

import (
	"fmt"
	"latihan/infrastructure/configuration"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

const (
	ErrFailedConnect string = "failed to initialize database"
	SlaveDB          string = "mygram"
)

// NewPostgre constructor
func NewPostgre(conf *configuration.AppConfig) *gorm.DB {

	connInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Postgre.User, conf.Postgre.Ps, conf.Postgre.Host, conf.Postgre.Port, conf.Postgre.Name)

	connSlaveInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.SlaveDB.User, conf.SlaveDB.Ps, conf.SlaveDB.Host, conf.SlaveDB.Port, conf.SlaveDB.Name)

	db, err := gorm.Open(postgres.Open(connInfo), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}
	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{postgres.Open(connSlaveInfo)},
		Policy:  dbresolver.RandomPolicy{},
	}, SlaveDB))
	if err != nil {
		logrus.Panic(err)
	}
	// db.Debug().AutoMigrate(uploadphoto.PhotoEntity{})
	return db
}

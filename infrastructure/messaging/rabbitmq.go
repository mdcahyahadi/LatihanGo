package messaging

import (
	"fmt"
	"latihan/infrastructure/configuration"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/sirupsen/logrus"
)

const (
	ErrFailedConnect string = "failed to initialize amqp"
)

func NewPublisher(conf *configuration.AppConfig) *amqp.Publisher {
	amqpConfig := amqp.NewDurableQueueConfig(fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.RabbitMQ.User, conf.RabbitMQ.Ps, conf.RabbitMQ.IP, conf.RabbitMQ.Port))
	publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		logrus.Errorf(ErrFailedConnect)
	}

	return publisher
}

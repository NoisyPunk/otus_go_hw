package queue

import (
	"net"

	senderconfig "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/sender_config"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type Consumer struct {
	RmqChannel    *amqp.Channel
	rmqConnection *amqp.Connection
}

func NewConsumer(log *zap.Logger, config *senderconfig.Config) (*Consumer, error) {
	url := "amqp://" + config.User + ":" + config.Password + "@" + net.JoinHostPort(config.Host, config.Port)
	connect, err := amqp.Dial(url)
	if err != nil {
		log.Error("consumer connection to rmq failed")
		return nil, errors.Wrap(err, "consimer connection to rmq failed")
	}

	channel, err := connect.Channel()
	if err != nil {
		log.Error("consumer channel creation failed")
		return nil, errors.Wrap(err, "consumer channel creation failed")
	}
	consumer := Consumer{
		RmqChannel:    channel,
		rmqConnection: connect,
	}
	return &consumer, nil
}

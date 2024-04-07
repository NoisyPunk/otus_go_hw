package queue

import (
	"context"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/sender_config"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"net"
)

type Consumer struct {
	RmqChannel    *amqp.Channel
	rmqConnection *amqp.Connection
}

func NewConsumer(ctx context.Context, config *sender_config.Config) (*Consumer, error) {
	l := logger.FromContext(ctx)
	url := "amqp://" + config.User + ":" + config.Password + "@" + net.JoinHostPort(config.Host, config.Port)
	connect, err := amqp.Dial(url)
	if err != nil {
		l.Error("consumer connection to rmq failed")
		return nil, errors.Wrap(err, "consimer connection to rmq failed")
	}

	channel, err := connect.Channel()
	if err != nil {
		l.Error("consumer channel creation failed")
		return nil, errors.Wrap(err, "consumer channel creation failed")
	}
	consumer := Consumer{
		RmqChannel:    channel,
		rmqConnection: connect,
	}
	return &consumer, nil
}

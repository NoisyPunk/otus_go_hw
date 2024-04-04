package queue

import (
	"context"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/scheduler_config"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"net"
)

type Producer struct {
	rmqConnection *amqp.Connection
	rmqChannel    *amqp.Channel
	rmqQueue      *amqp.Queue
}

func NewProducer(ctx context.Context, config *scheduler_config.Config) (*Producer, error) {
	l := logger.FromContext(ctx)
	url := "amqp://" + config.User + ":" + config.Password + "" + net.JoinHostPort(config.Host, config.Port)
	connect, err := amqp.Dial(url)
	if err != nil {
		l.Error("producer connection to rmq failed")
		return nil, errors.Wrap(err, "producer connection to rmq failed")
	}
	defer connect.Close()

	channel, err := connect.Channel()
	if err != nil {
		l.Error("producer channel creation failed")
		return nil, errors.Wrap(err, "producer channel creation failed")
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"CalendarQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		l.Error("producer queue creation failed")
		return nil, errors.Wrap(err, "producer queue creation failed")
	}
	producer := Producer{
		rmqConnection: connect,
		rmqChannel:    channel,
		rmqQueue:      &queue,
	}

	return &producer, nil
}

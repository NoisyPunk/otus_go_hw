package queue

import (
	"context"
	"net"

	schedulerConfig "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/scheduler_config"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type Producer struct {
	RmqChannel    *amqp.Channel
	rmqConnection *amqp.Connection
	rmqQueue      *amqp.Queue
}

func NewProducer(ctx context.Context, config *schedulerConfig.Config) (*Producer, error) {
	l := logger.FromContext(ctx)
	url := "amqp://" + config.User + ":" + config.Password + "@" + net.JoinHostPort(config.Host, config.Port)
	connect, err := amqp.Dial(url)
	if err != nil {
		l.Error("producer connection to rmq failed")
		return nil, errors.Wrap(err, "producer connection to rmq failed")
	}

	channel, err := connect.Channel()
	if err != nil {
		l.Error("producer channel creation failed")
		return nil, errors.Wrap(err, "producer channel creation failed")
	}

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
		RmqChannel:    channel,
		rmqConnection: connect,
		rmqQueue:      &queue,
	}

	return &producer, nil
}

package queue

import (
	"net"

	schedulerConfig "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/scheduler_config"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type Producer struct {
	RmqChannel    *amqp.Channel
	rmqConnection *amqp.Connection
	rmqQueue      *amqp.Queue
}

func NewProducer(log *zap.Logger, config *schedulerConfig.Config) (*Producer, error) {
	url := "amqp://" + config.RmqCreds.User + ":" + config.RmqCreds.Password + "@" + net.JoinHostPort(
		config.RmqCreds.Host, config.RmqCreds.Port)
	connect, err := amqp.Dial(url)
	if err != nil {
		log.Error("producer connection to rmq failed")
		return nil, errors.Wrap(err, "producer connection to rmq failed")
	}

	channel, err := connect.Channel()
	if err != nil {
		log.Error("producer channel creation failed")
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
		log.Error("producer queue creation failed")
		return nil, errors.Wrap(err, "producer queue creation failed")
	}
	producer := Producer{
		RmqChannel:    channel,
		rmqConnection: connect,
		rmqQueue:      &queue,
	}

	return &producer, nil
}

package queue

import (
	"context"
	"fmt"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/scheduler_config"
	"github.com/streadway/amqp"
	"net"
)

type Producer struct {
	rmqConnection *amqp.Connection
	rmqChannel    *amqp.Channel
	rmqQueue      *amqp.Queue
}

func NewProducer(ctx context.Context, config *scheduler_config.Config) (*Producer, error) {
	url := "amqp://" + config.User + ":" + config.Password + "@" + net.JoinHostPort(config.Host, config.Port)
	conn, err := amqp.Dial(url)
	//conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"CalendarQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("queue declared", q)

	err = ch.Publish(
		"",
		"CalendarQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("hello world"),
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Message published")
	return nil, nil
}

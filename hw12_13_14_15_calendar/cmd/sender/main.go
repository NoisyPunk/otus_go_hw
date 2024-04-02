package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {

	fmt.Println("consumer app")
	conn, err := amqp.Dial("amqp://user:pass@localhost:5672/")
	if err != nil {
		if err != nil {
			panic(err)
		}
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"CalendarQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	frv := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Println("recieved message: ", msg.Body)
		}
	}()

	fmt.Println("connected to rmq")
	fmt.Println("waotong for mesages")
	<-frv
}

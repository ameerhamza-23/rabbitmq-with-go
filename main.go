package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Successfully connected to rabbitmq instance")

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println("q : ", q)

	if err != nil {
		panic(err)
	}

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully published a message to queue")

}

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

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("received message %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully connected to rabbitmq instance")
	fmt.Printf("--- waiting for messages ---")

	<-forever

}

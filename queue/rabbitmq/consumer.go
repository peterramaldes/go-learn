package main

import (
	"github.com/streadway/amqp"
	"log"
)

const RABBIT_MQ_CONNECTION_URL = "amqp://guest:guest@localhost:5672/"

func main() {
	consume()
}

func consume() {
	conn, err := amqp.Dial(RABBIT_MQ_CONNECTION_URL)
	failOnError("Failed to connect to RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError("Failed to open a channel", err)
	defer ch.Close()

	q, err := ch.QueueDeclare("golang-queue", false, false, false, false, nil)
	failOnError("Failed to declare a queue", err)

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError("Failed to register a message", err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press Ctrl + C")
	<-forever
}

func failOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

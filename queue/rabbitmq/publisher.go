package main

import (
	"bufio"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

const RABBIT_MQ_CONNECTION_URL = "amqp://guest:guest@localhost:5672/"

func main() {
	publish()
}

func publish() {
	conn, err := amqp.Dial(RABBIT_MQ_CONNECTION_URL)
	failOnError("Failed to connect to RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError("Failed to open a channel", err)
	defer ch.Close()

	q, err := ch.QueueDeclare("golang-queue", false, false, false, false, nil)
	failOnError("Failed to declare a queue", err)

	// Let's catch the message from the terminal.
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What message do you want to send?")
	payload, _ := reader.ReadString('\n')

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(payload)})
	failOnError("Failed to publish a message", err)

	log.Printf(" [x] Congrats, sending message: %s", payload)
}

func failOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// This package listens to the NewBlockQueue,
// When a new block is published, it starts a
// worker, parses the trasnations, and pushes them
// to the trasnations queue.
package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	addr := "amqp://guest:guest@localhost:5672/"
	queueName := "NewBlockHash"

	// Create a connection to the RabbitMQ Server
	conn, err := amqp.Dial(addr)
	if err != nil {
		log.Fatalf("Unable to connect to RabbitMQ server: %s\n", err)
	}
	defer conn.Close()

	// Create a channel to send on
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to create a channel: %s\n", err)
	}
	defer ch.Close()

	// Define what Queue to publish on
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to create a Queue: %s\n", err)
	}

	hashes, err := ch.Consume(
		q.Name,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Unable to consume channel: %s", err)
	}

	forever := make(chan bool)
	go func() {
		for h := range hashes {
			fmt.Printf("Received Hash: %s\n", h.Body)
		}
	}()
	<-forever
}

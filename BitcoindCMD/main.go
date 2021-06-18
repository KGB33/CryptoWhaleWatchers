package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	// define subcommands
	startup := flag.NewFlagSet("startup", flag.ExitOnError)
	block := flag.NewFlagSet("block", flag.ExitOnError)

	// Add flags to subcommands
	startupMessagePtr := startup.String("message", "Bitcoind Started...", "Echos the provided string")
	blockHashPtr := block.String("hash", "", "The hash of the newly created Block")

	// Verify that a subcommand has been provided
	if len(os.Args) < 2 {
		fmt.Println("A subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "startup":
		if err := startup.Parse(os.Args[2:]); err != nil {
			fmt.Printf("Unable to parse provided flags: %s\n", err)
			os.Exit(1)
		}
		handleStartupCMD(*startupMessagePtr)
	case "block":
		if err := block.Parse(os.Args[2:]); err != nil {
			fmt.Printf("Unable to parse provided flags: %s\n", err)
			os.Exit(1)
		}
		handleStartupCMD(*startupMessagePtr)
		handleBlockCMD(*blockHashPtr)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

}

// Add the newly create block has to the event bus
// TODO: ^^
func handleBlockCMD(hash string) {
	// Create a connection to the RabbitMQ Server
	conn, err := amqp.Dial("amqp://guest:guest@message-broker:5672/")
	if err != nil {
		fmt.Printf("Unable to connect to RabbitMQ server: %s\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Create a channel to send on
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Failed to create a channel: %s\n", err)
		os.Exit(1)
	}
	defer ch.Close()

	// Define what Queue to publish on
	q, err := ch.QueueDeclare(
		"NewBlockHash", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		fmt.Printf("Failed to create a Queue: %s\n", err)
		os.Exit(1)
	}

	// Send the message!
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(hash),
		})
	if err != nil {
		fmt.Printf("Failed to publish hash Val: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n\n======== NEW BLOCK!! ==========\n%s\n\n", hash)
}

func handleStartupCMD(val string) {
	fmt.Println(val)
}

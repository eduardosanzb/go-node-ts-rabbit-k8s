package infrastructure

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

const (
	brokerUrl = "BROKER_URL"
	name      = "QUEUE_NAME"
)

var (
	serverUrl = os.Getenv(brokerUrl)
	queueName = os.Getenv(name)
)

type handler func([]byte) bool

func RabbitWorker(id int, fn handler) {
	connectRabbitMQ, err := amqp.Dial(serverUrl)
	if err != nil {
		log.Println("Error dialing the server")
		panic(err)
	}
	go func() {
		err := <-connectRabbitMQ.NotifyClose(make(chan *amqp.Error))
		log.Printf("closing the connection to rabbit: %v\n", <-connectRabbitMQ.NotifyClose(make(chan *amqp.Error)))
		if err != nil {
			panic(err)
		}
	}()

	channel, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	if err := channel.Qos(
		2,     // prefetch count
		0,     // prefetch size
		false, // global
	); err != nil {
		log.Println("Error Creating the output")
		panic(err)
	}

	queue, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Error declaring the queue")
		panic(err)
	}
	messagesChannel, err := channel.Consume(
		queue.Name, // queue name
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // arguments
	)
	if err != nil {
		log.Println("Failed to register a consumer")
		log.Println(err)
		panic(err)
	}
	for d := range messagesChannel {
		status := fn(d.Body)
		if err := d.Ack(!status); err != nil {
			log.Printf("Error acknowling message: %s", err)
		} else {
			log.Println("Acknowledged message")
		}
	}
	log.Printf(" [*]: Waiting for messages.")
}

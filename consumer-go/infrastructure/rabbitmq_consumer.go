package infrastructure

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

const brokerUrl = "BROKER_URL"

var serverUrl = os.Getenv(brokerUrl)

func CreateMessagesChannel(queueName string) <-chan amqp.Delivery {

	connectRabbitMQ, err := amqp.Dial(serverUrl)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	go func() {
		log.Printf("closing: %v\n", <-connectRabbitMQ.NotifyClose(make(chan *amqp.Error)))
	}()

	if err != nil {
		panic(err)
	}

	channel, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	channel.Qos(
		10,    // prefetch count
		0,     // prefetch size
		false, // global
	)

	messagesChannel, err := channel.Consume(
		queueName, // queue name
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // arguments
	)

	if err != nil {
		log.Println(err)
	}

	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	return messagesChannel
}

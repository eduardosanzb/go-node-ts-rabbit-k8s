package main

import (
	"log"
	"os"

	"github.com/eduardosanzb/consumer/usecases"

	"github.com/streadway/amqp"
)

func main() {

	forever := make(chan bool)
	for i := 1; i <= 1; i++ {
		go worker(i, usecases.CreateMessageWithSender)
	}
	<-forever

}

const brokerUrl = "BROKER_URL"
const name = "QUEUE_NAME"

var serverUrl = os.Getenv(brokerUrl)
var queueName = os.Getenv(name)

type handler func([]byte) bool

func worker(id int, fn handler) {
	connectRabbitMQ, err := amqp.Dial(serverUrl)
	if err != nil {
		panic(err)
	}
	go func() {
		err := <-connectRabbitMQ.NotifyClose(make(chan *amqp.Error))
		log.Printf("closing the connection to rabbit: %v\n", <-connectRabbitMQ.NotifyClose(make(chan *amqp.Error)))
		if err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

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
	}

	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	for d := range messagesChannel {
		status := fn(d.Body)
		log.Printf("status: %s\n", status)
		if err := d.Ack(!status); err != nil {
			log.Printf("Error acknowling message: %s", err)
		} else {
			log.Println("Acknowledged message")
		}
	}
	log.Printf(" [*]: Waiting for messages.")
}

package main

import (
	"github.com/eduardosanzb/consumer/infrastructure"
	"github.com/eduardosanzb/consumer/usecases"
)

func main() {

	forever := make(chan bool)
	for i := 1; i <= 1; i++ {
		go infrastructure.RabbitWorker(i, usecases.CreateMessageWithSender)
	}
	<-forever

}

package db

import (
	"sync"

	"github.com/streadway/amqp"
)

var (
	rabbitmqClient *amqp.Connection
	rabbitmqOnce   sync.Once
)

// @doc https://www.rabbitmq.com/tutorials/tutorial-one-go.html
func initRabbitmqClient() error {
	var err error
	rabbitmqClient, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	return nil
}

// Rabbitmq return rabbitmq client
func Rabbitmq() *amqp.Connection {
	if rabbitmqClient == nil {
		rabbitmqOnce.Do(func() {
			err := initRabbitmqClient()
			if err != nil {
				panic(err)
			}
		})
	}
	return rabbitmqClient
}

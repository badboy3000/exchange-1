package db

import (
	"sync"

	"github.com/streadway/amqp"
)

var (
	rabbitmqConnection *amqp.Connection
	rabbitmqChannel    *amqp.Channel
	rabbitmqOnce       sync.Once
)

// @doc https://www.rabbitmq.com/tutorials/tutorial-two-go.html
func initRabbitmqConnection() error {
	var err error
	rabbitmqConnection, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	return nil
}

func initRabbitmqChannel() error {
	var err error
	rabbitmqChannel, err = rabbitmqConnection.Channel()
	if err != nil {
		return err
	}
	return nil
}

// DeclareMatchingWorkQueue ...
func DeclareMatchingWorkQueue() amqp.Queue {
	rabbitmqQ, err := rabbitmqChannel.QueueDeclare(
		"exchange.matching.work.queue", // name
		true,                           // durable
		false,                          // delete when unused
		false,                          // exclusive
		false,                          // no-wait
		nil,                            // arguments
	)
	if err != nil {
		panic(err)
	}
	return rabbitmqQ
}

// RabbitmqChannel return rabbitmq channel
func RabbitmqChannel() *amqp.Channel {
	if rabbitmqConnection == nil {
		rabbitmqOnce.Do(func() {
			err := initRabbitmqConnection()
			if err != nil {
				panic(err)
			}
			err = initRabbitmqChannel()
			if err != nil {
				panic(err)
			}
		})
	}
	return rabbitmqChannel
}

// DeclareExchange ...
func DeclareExchange(xName, xType string) {
	err := rabbitmqChannel.ExchangeDeclare(
		xName, // name
		xType, // type direct
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		panic(err)
	}
}

// DeclareQueue ...
func DeclareQueue(name string) amqp.Queue {
	q, err := rabbitmqChannel.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		panic(err)
	}
	return q
}

// Bind ...
func Bind(name, routingKey, xName string) {
	err := rabbitmqChannel.QueueBind(
		name,       // queue name
		routingKey, // routing key
		xName,      // exchange name
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		panic(err)
	}
}

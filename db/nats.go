package db

import (
	"sync"

	"github.com/nats-io/go-nats"
)

var (
	natsClient *nats.Conn
	natsOnce   sync.Once
)

// @doc https://github.com/nats-io/nats.go
func initNatsClient() error {
	var err error
	natsClient, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		return err
	}
	return nil
}

// Nats return nats client
func Nats() *nats.Conn {
	if natsClient == nil {
		natsOnce.Do(func() {
			err := initNatsClient()
			if err != nil {
				panic(err)
			}
		})
	}
	return natsClient
}

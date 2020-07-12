package rpc 

import (
	"os"
	"errors"

	"github.com/nats-io/nats.go"
)

func natsConnect() (*nats.Conn, error) {
	uri := os.Getenv("NATS_URI")
	if len(uri) == 0 {
		return nil, errors.New("Invalid env: NATS_URI")
	}
	nc, err := nats.Connect(uri)
	if err != nil {
		return nil, err 
	}
	return nc, nil 
}
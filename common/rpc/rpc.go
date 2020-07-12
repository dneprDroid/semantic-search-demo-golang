package rpc

import (
	"time"
	"errors"
	"encoding/json"

	"common/log"

	"github.com/nats-io/nats.go"
)

type Rpc struct {
	nc *nats.Conn
}

type ListenHandler func(request TextRequest) (*EmbeddingResponse, error)

func NewRpc() *Rpc {
	return &Rpc{}
}

func (self *Rpc) Connect() error {
	nc, err := natsConnect()
	if err != nil {
		return err 
	}
	self.nc = nc 
	return nil 
}

func (self *Rpc) Listen(subject string, handler ListenHandler) error {
	_, err := self.nc.Subscribe(subject, func(m *nats.Msg) {
		respWithErr := func (err error) {
			log.Printf("[Listen] err: '%v'", err)
			m.Respond(nil)
		}
		var request TextRequest
		if err := json.Unmarshal(m.Data, &request); err != nil {
			respWithErr(err)
			return 
		}
		resp, err := handler(request)
		if err != nil {
			respWithErr(err)
			return 
		}
		if resp == nil {
			log.Printf("[Listen] resp nil")
			respWithErr(nil)
			return 
		}
		respBytes, err := json.Marshal(*resp)
		if err != nil {
			respWithErr(err)
			return 
		}
		m.Respond(respBytes)
	})
	return err 
}

func (self *Rpc) Request(subject string, request TextRequest) (*EmbeddingResponse, error) {
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err 
	}
	msg, err := self.nc.Request(subject, reqBytes, 10 * time.Millisecond)
	if err != nil {
		return nil, err
	}
	if len(msg.Data) == 0 {
		return nil, errors.New("Invalid message")
	}
	resp := new(EmbeddingResponse)
	if err := json.Unmarshal(msg.Data, resp); err != nil {
		return nil, err 
	}
	return resp, nil 
}

func (self *Rpc) Release() {
	if self.nc == nil {
		return 
	}
	self.nc.Flush()
	self.nc.Close()
}
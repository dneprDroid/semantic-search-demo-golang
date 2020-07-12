package main 

import (
	"os"
	"os/signal"
	"syscall"

	"common/rpc"
	"common/log"
)

func waitServ() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	for {
		// Wait for signal
		<-c
		log.Printf("<caught signal - draining>")
	}
	log.Printf("Stopped!")
}

func main() {
	rpcEngine := rpc.NewRpc()

	if err := rpcEngine.Connect(); err != nil {
		panic(err)
	}
	defer rpcEngine.Release()

	err := rpcEngine.Listen("embedding", func (request rpc.TextRequest) (*rpc.EmbeddingResponse, error) {

		return nil, nil 
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Running text2embeddings ...")
	waitServ()
}
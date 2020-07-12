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
	tk := NewTokenizer()

	if err := rpcEngine.Connect(); err != nil {
		panic(err)
	}
	defer rpcEngine.Release()

	err := rpcEngine.Listen("embedding", func (request rpc.TextRequest) (*rpc.EmbeddingResponse, error) {
		log.Printf("got new request: `%v`", request)
		vec, err := tk.Encode(request.Content)
		if err != nil {
			log.Printf("tokenizer err: `%v`", err)
			return nil, err
		}
		log.Printf("got new response: `%v` => `%v`", request.Content, vec)
		resp := &rpc.EmbeddingResponse{
			Vec: vec,
		}
		return resp, nil 
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Running text2embeddings ...")
	waitServ()
}
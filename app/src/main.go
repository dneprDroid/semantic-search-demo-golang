package main

import (
	"fmt"
	"net/http"
	"os"

	servhandlers "appserver/handlers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	server := &http.Server{
		Addr: ":" + os.Getenv("SERVER_PORT"),
	}

	sh := servhandlers.NewServer()

	http.Handle("/", r)

	fmt.Printf("starting server :: %#v \n", server.Addr)

	http.ListenAndServe(server.Addr, handlers.LoggingHandler(os.Stdout, r))
}
package main

import (
	"fmt"
	"net/http"
	"os"

	servhandlers "appserver/handlers"
	"appserver/handlers/posts"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	server := &http.Server{
		Addr: ":" + os.Getenv("SERVER_PORT"),
	}

	sh := servhandlers.NewServer()

	r.HandleFunc("/post", sh.Handler(posts.AddPost))
	r.HandleFunc("/post/search", sh.Handler(posts.FindPost))

	http.Handle("/", r)

	fmt.Printf("starting server :: %#v \n", server.Addr)

	http.ListenAndServe(server.Addr, handlers.LoggingHandler(os.Stdout, r))
}
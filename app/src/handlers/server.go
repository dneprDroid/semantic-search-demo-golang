package handlers 

import (
	"fmt"
	"net/http"

	"common/log"

	"appserver/database"
	"appserver/env"
	"appserver/util"
)

type ServerHandler func(env.Env, http.ResponseWriter, *http.Request)

type Server struct {
	env env.Env
}

func NewServer() *Server {
	return &Server{
		env: env.Env{
			Db: database.NewPool(),
		},
	} 
}

func (self* Server) Handler(h ServerHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered. Error: `%v`", err)
				util.RespErrorStr(w, fmt.Sprintf("%v", err))
			}
		}()
		h(self.env, w, req)
	}
}
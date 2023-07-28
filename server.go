package MyToDoMadeOnGo

import (
	"context"
	"net/http"
	"time"
)

type MyServer struct {
	httpServer *http.Server
}

func (s *MyServer) Run(port string) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *MyServer) Shutdown(cxt context.Context) error {
	return s.httpServer.Shutdown(cxt)
}

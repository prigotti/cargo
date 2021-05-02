package server

import (
	"context"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	server   *http.Server
	listener net.Listener
}

func SetupHTTPServer(ctx context.Context, router *mux.Router, bindAddress string) (*HTTPServer, error) {
	var err error
	httpServer := &HTTPServer{}
	lc := net.ListenConfig{}

	httpServer.listener, err = lc.Listen(ctx, "tcp", bindAddress)
	if err != nil {
		return nil, err
	}

	server := http.Server{
		Addr:    httpServer.listener.Addr().String(),
		Handler: router,
	}

	go func() {
		server.Serve(httpServer.listener)
	}()

	return httpServer, nil
}

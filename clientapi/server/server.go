package server

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/prigotti/cargo/clientapi/application"
	"github.com/prigotti/cargo/clientapi/infrastructure/adapter/http/handler"
	"google.golang.org/grpc"
)

// Server holds all dependencies for this microservice
type Server struct {
	configuration *Configuration
	httpServer    *HTTPServer
	grpcClient    *grpc.ClientConn
	portService   application.PortService
}

// NewServer is the Server factory
func NewServer(ctx context.Context, configuration *Configuration) (*Server, error) {
	s := &Server{configuration: configuration}
	var err error
	router := mux.NewRouter()

	s.grpcClient, err = SetupGRPCClient(s.configuration.GRPCServerAddress)
	if err != nil {
		return nil, err
	}

	s.portService = application.NewPortService(s.grpcClient)

	handler.NewPortHTTPHandler(router, s.portService)

	s.httpServer, err = SetupHTTPServer(ctx, router, s.configuration.HTTPServerBindAddress)
	if err != nil {
		return nil, err
	}

	return s, nil
}

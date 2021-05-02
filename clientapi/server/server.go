package server

import (
	"context"

	"github.com/prigotti/cargo/clientapi/application"
	"google.golang.org/grpc"
)

// Server holds all dependencies for this microservice
type Server struct {
	configuration *Configuration
	grpcClient    *grpc.ClientConn
	portService   application.PortService
}

// NewServer is the Server factory
func NewServer(ctx context.Context, configuration *Configuration) (*Server, error) {
	s := &Server{configuration: configuration}
	var err error

	s.grpcClient, err = SetupGRPCClient(s.configuration.GRPCServerAddress)
	if err != nil {
		return nil, err
	}

	s.portService = application.NewPortService(s.grpcClient)

	return s, nil
}

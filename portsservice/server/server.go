package server

import (
	"context"

	"github.com/prigotti/cargo/portsservice/application"
	"github.com/prigotti/cargo/portsservice/domain"
	"github.com/prigotti/cargo/portsservice/infrastructure/adapter/repository"
	"google.golang.org/grpc"
)

// Server holds all dependencies for this microservice
type Server struct {
	configuration  *Configuration
	grpcServer     *GRPCServer
	portRepository domain.PortRepository
	portService    application.PortService
}

// NewServer is the Server factory
func NewServer(ctx context.Context, configuration *Configuration) (*Server, error) {
	s := &Server{configuration: configuration}
	var err error
	gs := grpc.NewServer()

	s.portRepository = repository.NewMongoDBPortRepository()
	if err != nil {
		return nil, err
	}

	s.portService = application.NewPortService(gs, s.portRepository)

	s.grpcServer, err = SetupGRPCServer(ctx, gs, s.configuration.GRPCServerBindAddress)
	if err != nil {
		return nil, err
	}

	return s, nil
}

package server

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

// GRPCServer holds the gRPC server and its listener
type GRPCServer struct {
	server   *grpc.Server
	listener net.Listener
}

// SetupGRPCServer configures and launches a gRPC server in its own goroutine
func SetupGRPCServer(ctx context.Context, server *grpc.Server, bindAddress string) (*GRPCServer, error) {
	var err error
	grpcServer := &GRPCServer{server: server}
	lc := net.ListenConfig{}

	grpcServer.listener, err = lc.Listen(ctx, "tcp", bindAddress)
	if err != nil {
		return nil, err
	}

	go func() {
		server.Serve(grpcServer.listener)
	}()

	return grpcServer, nil
}

package server

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	server   *grpc.Server
	listener net.Listener
}

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

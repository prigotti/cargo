package server

import (
	"fmt"

	"google.golang.org/grpc"
)

// SetupGRPCClient Configures and returns a gRPC client for the
// given server address.
func SetupGRPCClient(serverAddress string) (*grpc.ClientConn, error) {
	fmt.Println(serverAddress)
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

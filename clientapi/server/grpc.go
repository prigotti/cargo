package server

import (
	"fmt"

	"google.golang.org/grpc"
)

func SetupGRPCClient(serverAddress string) (*grpc.ClientConn, error) {
	fmt.Println(serverAddress)
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

package server

import "context"

// Server holds all dependencies for this microservice
type Server struct {
	configuration *Configuration
}

// NewServer is the Server factory
func NewServer(ctx context.Context, configuration *Configuration) (*Server, error) {
	return &Server{configuration}, nil
}

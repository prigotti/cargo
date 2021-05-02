package server

import (
	"os"

	"github.com/imdario/mergo"
	"github.com/joho/godotenv"
)

// Default values and environment variable map keys
const (
	defaultJSONPath          = "./ports.json"
	defaultGRPCServerAddress = "localhost:9002"

	jsonPathKey       = "JSON_PATH"
	grpcServerAddress = "GRPC_SERVER_ADDRESS"
)

// Configuration holds overall server configuration data
type Configuration struct {
	JSONPath          string
	GRPCServerAddress string
}

// DefaultConfiguration retrieves default Configuration data
func DefaultConfiguration() *Configuration {
	return &Configuration{
		JSONPath:          defaultJSONPath,
		GRPCServerAddress: defaultGRPCServerAddress,
	}
}

// GetEnvironmentConfiguration retrieves environment configuration
// (also from a .env file, if it exists)
func GetEnvironmentConfiguration() *Configuration {
	_ = godotenv.Load()

	return &Configuration{
		JSONPath:          os.Getenv(jsonPathKey),
		GRPCServerAddress: os.Getenv(grpcServerAddress),
	}
}

// Merge combines two configurations, giving precedence to variables
// in the dst instance
func (c *Configuration) Merge(dst *Configuration) (*Configuration, error) {
	err := mergo.Merge(dst, c)
	if err != nil {
		return nil, err
	}

	return dst, nil
}

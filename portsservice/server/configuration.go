package server

import (
	"os"

	"github.com/imdario/mergo"
	"github.com/joho/godotenv"
)

// Default values and environment variable map keys
const (
	defaultGRPCServerBindAddress = "0.0.0.0:9002"
	defaultDatabaseName          = "PortServiceDatabase"
	defaultDatabaseURI           = "mongodb://localhost:27017"
	defaultDatabaseUser          = "admin"
	defaultDatabasePassword      = "secret"

	databaseURIKey = "DATABASE_URI"
)

// Configuration holds overall server configuration data
type Configuration struct {
	GRPCServerBindAddress string
	DatabaseURI           string
	DatabaseName          string
	DatabaseUser          string
	DatabasePassword      string
}

// DefaultConfiguration retrieves default Configuration data
func DefaultConfiguration() *Configuration {
	return &Configuration{
		GRPCServerBindAddress: defaultGRPCServerBindAddress,
		DatabaseName:          defaultDatabaseName,
		DatabaseURI:           defaultDatabaseURI,
		DatabaseUser:          defaultDatabaseUser,
		DatabasePassword:      defaultDatabasePassword,
	}
}

// GetEnvironmentConfiguration retrieves environment configuration
// (also from a .env file, if it exists)
func GetEnvironmentConfiguration() *Configuration {
	_ = godotenv.Load()

	return &Configuration{
		DatabaseURI: os.Getenv(databaseURIKey),
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

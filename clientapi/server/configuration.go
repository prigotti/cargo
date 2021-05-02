package server

import (
	"github.com/imdario/mergo"
	"github.com/joho/godotenv"
)

const ()

// Configuration holds overall server configuration data
type Configuration struct {
}

// DefaultConfiguration retrieves default Configuration data
func DefaultConfiguration() *Configuration {
	return &Configuration{}
}

// GetEnvironmentConfiguration retrieves environment configuration
// (also from a .env file, if it exists)
func GetEnvironmentConfiguration() *Configuration {
	_ = godotenv.Load()

	return &Configuration{}
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

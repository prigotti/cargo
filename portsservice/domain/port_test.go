package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPort(t *testing.T) {
	// We only have Coordinate validation, so that
	// should be enough for testing
	tests := []struct {
		id          string
		name        string
		city        string
		country     string
		alias       []string
		regions     []string
		coordinates []float32
		province    string
		timezone    string
		unlocs      []string
		code        string
		withError   bool
	}{
		{
			id:          "ABCDE",
			name:        "",
			city:        "",
			country:     "",
			alias:       []string{},
			regions:     []string{},
			coordinates: []float32{3.23, 9.46, 180.1},
			province:    "",
			timezone:    "",
			unlocs:      []string{},
			code:        "",
			withError:   true,
		},
		{
			id:          "ABCDE",
			name:        "",
			city:        "",
			country:     "",
			alias:       []string{},
			regions:     []string{},
			coordinates: []float32{-190.23, 9.46},
			province:    "",
			timezone:    "",
			unlocs:      []string{},
			code:        "",
			withError:   true,
		},
		{
			id:          "ABCDE",
			name:        "",
			city:        "",
			country:     "",
			alias:       []string{},
			regions:     []string{},
			coordinates: []float32{3.23, 9.46},
			province:    "",
			timezone:    "",
			unlocs:      []string{},
			code:        "",
			withError:   false,
		},
	}

	for _, test := range tests {
		_, err := NewPort(
			test.id,
			test.name,
			test.city,
			test.country,
			test.alias,
			test.regions,
			test.coordinates,
			test.province,
			test.timezone,
			test.unlocs,
			test.code,
		)
		if test.withError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

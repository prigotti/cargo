package domain

import (
	"context"
	"math"
)

const (
	// Simple invariants added for illustration
	CoordinateNumber = 2
	LongitudeIndex   = 0
	LatitudeIndex    = 1
	MaxAbsLongitude  = 180.0
	MaxAbsLatitude   = 90.0
)

// Port is the aggregate (clearly root and only entity) of
// the Port domain model.
//
// Encapsulation and "defensive code" is not applied as it
// can make rehydration harder on the repository implementation
// that lives in another package, besides adding a lot of
// complexity without fully achieving OO standards.
// It is then required that state changes be made only through
// the aggregate's methods so that invariants are respected,
// even though there's little behavior and validation happening
// right now.
//
// Value objects were not designed for simplicity, but could make
// many operations easier after implemented.
type Port struct {
	ID          string    `json:"id" bson:"id"`
	Name        string    `json:"name" bson:"name"`
	City        string    `json:"city" bson:"city"`
	Country     string    `json:"country" bson:"country"`
	Alias       []string  `json:"alias" bson:"alias"`
	Regions     []string  `json:"regions" bson:"regions"`
	Coordinates []float32 `json:"coordinates" bson:"coordinates"`
	Province    string    `json:"province" bson:"province"`
	Timezone    string    `json:"timezone" bson:"timezone"`
	Unlocs      []string  `json:"unlocs" bson:"unlocs"`
	Code        string    `json:"code" bson:"code"`
}

// NewPort is the Port factory and applies validation and
// invariants.
func NewPort(
	id string,
	name string,
	city string,
	country string,
	alias []string,
	regions []string,
	coordinates []float32,
	province string,
	timezone string,
	unlocs []string,
	code string,
) (*Port, error) {
	if !validateCoordinates(coordinates) {
		return nil, ErrInvalidAggregateData
	}

	// Omitted other validations (also due to lack of domain knowledge)

	return &Port{
		ID:          id,
		Name:        name,
		City:        city,
		Country:     country,
		Alias:       alias,
		Regions:     regions,
		Coordinates: coordinates,
		Province:    province,
		Timezone:    timezone,
		Unlocs:      unlocs,
		Code:        code,
	}, nil
}

// Update looks like the Port factory at the moment, but
// as we gather more domain knowledge, it could be broken
// into more specific methods that could possibly generate
// domain events, which would make this a much richer model.
func (p *Port) Update(
	name string,
	city string,
	country string,
	alias []string,
	regions []string,
	coordinates []float32,
	province string,
	timezone string,
	unlocs []string,
	code string,
) error {
	if !validateCoordinates(coordinates) {
		return ErrInvalidAggregateData
	}

	// Omitted other validations (also due to lack of domain knowledge)

	p.Name = name
	p.City = city
	p.Country = country
	p.Alias = alias
	p.Regions = regions
	p.Coordinates = coordinates
	p.Province = province
	p.Timezone = timezone
	p.Unlocs = unlocs
	p.Code = code

	return nil
}

// Would be easier if designed as a method of a Coordinate
// value object.
func validateCoordinates(c []float32) bool {
	if cLen := len(c); cLen > CoordinateNumber || cLen == 0 {
		return false
	}

	if math.Abs(float64(c[LongitudeIndex])) > MaxAbsLongitude {
		return false
	}

	if math.Abs(float64(c[LatitudeIndex])) > MaxAbsLatitude {
		return false
	}

	return true
}

// PortRepository is the Port aggregate repository interface
// that'll be implemented in the infrastructure layer.
type PortRepository interface {
	Save(ctx context.Context, port *Port) error
	GetWithID(ctx context.Context, id string) (*Port, error)
}

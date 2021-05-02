package domain

import "errors"

const (
	errAggregateNotFound    = "aggregate not found"
	errInvalidAggregateData = "aggregate data is invalid"
)

var (
	ErrAggregateNotFound    = errors.New(errAggregateNotFound)
	ErrInvalidAggregateData = errors.New(errInvalidAggregateData)
)

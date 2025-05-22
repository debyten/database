package database

import "context"

// Provider must be implemented by database providers.
type Provider[T any] interface {
	// Conn should return a db connection.
	Conn(ctx ...context.Context) T
}
package database

import "io"

// Conn can be used to execute common operations on the database.
type Conn[T any] interface {
	Provider[T]
	Tx[T]

	io.Closer
}
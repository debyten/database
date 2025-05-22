package database

import (
	"context"
)

// TxDB exposes Provider and Transaction functions.
type TxDB[T any] interface {
	Provider[T]
	Transaction
	// Ctx must be used after Begin() to propagate the new context with
	// the TxDB instance
	Ctx() context.Context
}

// Transaction represents the operations on a database transaction.
type Transaction interface {
	// Commit the transaction. If the transaction was not started in the current context, it will be ignored.
	// If any errors are provided in the e parameter, the transaction will be rolled back (but only for the parent context).
	Commit(e ...error) error
	Rollback(err error) error
}

type Tx[T any] interface {
	// Begin starts a transaction. If the transaction is already started,
	// it continues from the underlying transaction context.
	Begin(ctx context.Context) TxDB[T]
}

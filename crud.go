package database

import "context"

// Crud is an interface that defines a set of methods for performing create, read, update, and delete
// operations on an object collection of type T identified by a unique ID of type ID. The ID type
// must be comparable, meaning that it must support the == and != operators.
type Crud[T any, ID comparable] interface {
	// Create inserts a new object of type T into the collection.
	Create(ctx context.Context, entity *T) error
	CreateMany(ctx context.Context, entities *[]T) error
	// Save updates an existing object of type T in the collection, replacing its current values with
	// the values of the provided object.
	Save(ctx context.Context, entity *T) error
	// SaveMany updates multiple objects of type T in the collection at once.
	SaveMany(ctx context.Context, entities *[]T) error
	// Delete entities by query. If no query is provided, an error is returned.
	Delete(ctx context.Context, query ...any) error
	// Update updates an existing object of type T in the collection, modifying its current values with
	// the values provided in the update operation.
	Update(ctx context.Context, entity *T) error
	// Count returns the total number of objects in the collection that match the specified query. If no
	// query is provided, it returns the total number of objects in the collection.
	Count(ctx context.Context, query ...any) (int64, error)
	// FindPage accept query which if not empty must contain `column_name: value`.
	FindPage(ctx context.Context, offset, size int, query ...any) ([]T, error)
	FindByID(ctx context.Context, id ID) (*T, error)
	FindOneBy(ctx context.Context, query ...any) (*T, error)
	// FindBy returns all the objects in the collection that match the specified query. If no query is
	// provided, it returns all the objects in the collection.
	FindBy(ctx context.Context, query ...any) ([]T, error)
	ExistsByID(ctx context.Context, id ID) (bool, error)
	// ExistsBy returns true if at least one object in the collection matches the specified query. If no
	// query is provided, it returns an error.
	ExistsBy(ctx context.Context, query ...any) (bool, error)
	// MustExistByID return an error that identifies a database error or the non-existence of the
	// searched entity (by id field).
	//
	// When the entity doesn't exist, an apierr.NotFound error is returned.
	MustExistByID(ctx context.Context, id ID) error
	// MustExistBy ensures that at least one entity matching the given query exists, returning an error if none are found.
	// If no query is provided, it returns an error.
	MustExistBy(ctx context.Context, query ...any) error
}

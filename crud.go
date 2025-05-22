package database

import "context"

// Crud is an interface that defines a set of methods for performing create, read, update, and delete
// operations on an object collection of type T identified by a unique ID of type ID. The ID type
// must be comparable, meaning that it must support the == and != operators.
type Crud[T any, ID comparable] interface {
	// Create inserts a new object of type T into the collection.
	Create(ctx context.Context, entity *T) error
	// CreateMany inserts multiple objects of type T into the collection at once.
	CreateMany(ctx context.Context, entities *[]T) error
	// Save updates an existing object of type T in the collection, replacing its current values with
	// the values of the provided object.
	Save(ctx context.Context, entity *T) error
	// SaveMany updates multiple objects of type T in the collection at once.
	SaveMany(ctx context.Context, entities *[]T) error
	// Delete removes an object of type T from the collection.
	Delete(ctx context.Context, entity *T) error
	// DeleteByIDs removes multiple objects from the collection at once, identified by their unique IDs.
	DeleteByIDs(ctx context.Context, ids []ID) error
	// Update updates an existing object of type T in the collection, modifying its current values with
	// the values provided in the update operation.
	Update(ctx context.Context, entity *T) error
	// Count returns the total number of objects in the collection that match the specified query. If no
	// query is provided, it returns the total number of objects in the collection.
	Count(ctx context.Context, query ...map[string]any) (int64, error)
	// CountByIDs returns the number of objects in the collection identified by the provided list of IDs.
	CountByIDs(ctx context.Context, ids []ID) (int64, error)
	// FindAll returns a slice containing all the objects in the collection.
	FindAll(ctx context.Context) ([]T, error)
	// FindPage accept query which if not empty must contain `column_name: value`.
	FindPage(ctx context.Context, offset, size int, query ...map[string]any) ([]T, error)
	FindByID(ctx context.Context, id ID) (*T, error)
	FindByIDs(ctx context.Context, id []ID) ([]T, error)
	// FindByIDAndCreatedBy return a single element filtered by id and `created_by` field.
	FindByIDAndCreatedBy(ctx context.Context, id ID, createdBy string) (*T, error)
	FindByCreatedBy(ctx context.Context, createdBy string) ([]T, error)

	ExistsByID(ctx context.Context, id ID) (bool, error)
	ExistsByIDAndCreatedBy(ctx context.Context, id ID, createdBy string) (bool, error)
	ExistsByProperty(ctx context.Context, propertyValue any, property string) (bool, error)
	// MustExistByID return an error that identifies a database error or the non-existence of the
	// searched entity (by id field).
	//
	// When the entity doesn't exist an apierr.NotFound error is returned.
	MustExistByID(ctx context.Context, id ID) error
	// MustExistByIDAndCreatedBy return an error that identifies a database error or the non-existence of the
	// searched entity (by id and created_by fields).
	//
	// When the entity doesn't exist an apierr.NotFound error is returned.
	MustExistByIDAndCreatedBy(ctx context.Context, id ID, createdBy string) error

	QueryMany(ctx context.Context, query ...any) ([]T, error)
}

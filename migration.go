package database

import (
	"context"
	"errors"
	"io/fs"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type MigrateDriver func(ctx context.Context) (database.Driver, error)

// NewMigrator returns a new unstarted migration engine.
//
//	Example:
//	m, err := NewMigrator(fs, dbCtx, dbConfig)
//	// handle err
//	if err := m.Up(); err != nil {
//	  // handle...
//	}
func NewMigrator(ctx context.Context, fs fs.FS, dr MigrateDriver, path string, execMigrations bool) (*migrate.Migrate, error) {
	m, err := dr(ctx)
	if err != nil {
		return nil, err
	}
	d, err := iofs.New(fs, path)
	if err != nil {
		return nil, err
	}
	inst, err := migrate.NewWithInstance("iofs", d, "migrations", m)
	if err != nil {
		return nil, err
	}
	if !execMigrations {
		return inst, nil
	}
	err = inst.Up()
	if err == nil {
		return inst, nil
	}
	if !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}
	return inst, nil
}

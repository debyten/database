package dbconf

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type SQLite struct {
	// Path points to the file that holds the sqlite db.
	// When providing `:memory:` sqlite will run in memory.
	Path string `default:"../sqlite.db"`
	// UpgradePath is the path to the database's upgrade script.
	UpgradePath string `default:"migrations"`
}

func NewSQLiteInMemory(migrationPath string) Datasource {
	return SQLite{Path: ":memory:", UpgradePath: migrationPath}
}

func NewSQLite() Datasource {
	var s SQLite
	envconfig.MustProcess("SQLITE", &s)
	return s
}

func (SQLite) DBName() string {
	return "sqlite"
}

// ConnURL returns the connection URL for the MySQL server.
func (s SQLite) ConnURL() string {
	return s.Path
}

func (s SQLite) ConnURLWithPrefix() string {
	return fmt.Sprintf("sqlite://%s", s.ConnURL())
}

// GetUpgradePath returns the path to the database's upgrade script.
func (s SQLite) GetUpgradePath() string {
	return s.UpgradePath
}

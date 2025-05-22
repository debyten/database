package dbconf

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

const PostgresFmt = "host=%s user=%s password=%s port=%d dbname=%s sslmode=disable"

// Postgres is a struct for storing configuration information for connecting to a Postgres database.
type Postgres struct {
	// User is the username for authenticating with the Postgres server.
	User string `default:"root"`
	// Password is the password for authenticating with the Postgres server.
	Password string `default:"root"`
	// Host is the hostname or IP address of the Postgres server.
	Host string `default:"127.0.0.1"`
	// Port is the port number of the Postgres server.
	Port uint `default:"5432"`
	// DB is the name of the database to connect to.
	DB string `default:"DEFAULT_DB"`
	// UpgradePath is the path to the database's upgrade script.
	UpgradePath string `default:"db"`
}

func (m Postgres) DBName() string {
	return m.DB
}

// ConnURL returns the connection URL for the Postgres server.
func (m Postgres) ConnURL() string {
	return fmt.Sprintf(PostgresFmt, m.Host, m.User, m.Password, m.Port, m.DB)
}

// GetUpgradePath returns the path to the database's upgrade script.
func (m Postgres) GetUpgradePath() string {
	return m.UpgradePath
}

// NewPostgres returns a new Postgres struct with configuration information read from the environment variables.
// If the DB field is set to "DEFAULT_DB", it will be replaced with the defaultDB parameter.
func NewPostgres(defaultDB string) Postgres {
	var m Postgres
	if err := envconfig.Process("SQL", &m); err != nil {
		panic(err)
	}
	if m.DB == "DEFAULT_DB" {
		m.DB = defaultDB
	}
	return m
}

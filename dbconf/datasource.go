package dbconf

// Datasource is an interface for managing a database connection and access to its properties.
type Datasource interface {
	// ConnURL returns the connection URL for the database.
	ConnURL() string

	// ConnURLWithPrefix returns the connection URL for the database with the connection prefix added (e.g. mysql://).
	ConnURLWithPrefix() string
	// DBName returns the name of the database.
	DBName() string
	// GetUpgradePath returns the path to the database's upgrade script.
	GetUpgradePath() string
}

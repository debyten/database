package dbconf

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

const mysqlFmt = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true"

// Mysql is a struct for storing configuration information for connecting to a MySQL database.
type Mysql struct {
	// User is the username for authenticating with the MySQL server.
	User string `default:"root"`
	// Password is the password for authenticating with the MySQL server.
	Password string `default:"root"`
	// Host is the hostname or IP address of the MySQL server.
	Host string `default:"127.0.0.1"`
	// Port is the port number of the MySQL server.
	Port uint `default:"3306"`
	// DB is the name of the database to connect to.
	DB string `default:"DEFAULT_DB"`
	// UpgradePath is the path to the database's upgrade script.
	UpgradePath string `default:"db"`
}

func (m Mysql) DBName() string {
	return m.DB
}

// ConnURL returns the connection URL for the MySQL server.
func (m Mysql) ConnURL() string {
	return fmt.Sprintf(mysqlFmt, m.User, m.Password, m.Host, m.Port, m.DB)
}

// GetUpgradePath returns the path to the database's upgrade script.
func (m Mysql) GetUpgradePath() string {
	return m.UpgradePath
}

// NewMysql returns a new Mysql struct with configuration information read from the environment variables.
// If the DB field is set to "DEFAULT_DB", it will be replaced with the defaultDB parameter.
func NewMysql(defaultDB string) Mysql {
	var m Mysql
	if err := envconfig.Process("SQL", &m); err != nil {
		panic(err)
	}
	if m.DB == "DEFAULT_DB" {
		m.DB = defaultDB
	}
	return m
}

package dbconf

import (
	"fmt"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

const mysqlFmt = "%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local&multiStatements=true&%s"

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
	UpgradePath string `default:"migrations"`

	params map[string]string
}

func (m Mysql) DBName() string {
	return m.DB
}

// ConnURL returns the connection URL for the MySQL server.
func (m Mysql) ConnURL() string {
	additionalParams := ""
	for k, v := range m.params {
		additionalParams += fmt.Sprintf("%s=%s&", k, v)
	}
	additionalParams = strings.TrimRight(additionalParams, "&")
	return fmt.Sprintf(mysqlFmt, m.User, m.Password, m.Host, m.Port, m.DB, additionalParams)
}

// GetUpgradePath returns the path to the database's upgrade script.
func (m Mysql) GetUpgradePath() string {
	return m.UpgradePath
}

func (m Mysql) ConnURLWithPrefix() string {
	return fmt.Sprintf("mysql://%s", m.ConnURL())
}

// NewMysql returns a new Mysql struct with configuration information read from the environment variables.
// If the DB field is set to "DEFAULT_DB", it will be replaced with the defaultDB parameter.
//
// params is a map of additional parameters to be passed to the database driver.
func NewMysql(defaultDB string, params ...map[string]string) Mysql {
	var m Mysql
	if err := envconfig.Process("SQL", &m); err != nil {
		panic(err)
	}
	if m.DB == "DEFAULT_DB" {
		m.DB = defaultDB
	}
	if len(params) > 0 {
		m.params = params[0]
	}
	return m
}

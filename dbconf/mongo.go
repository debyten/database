package dbconf

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

const (
	simpleDBConnFmt = "mongodb://%s:%d/?replicaSet=%s&directConnection=%t"
	fullDBConnFmt   = "mongodb://%s:%s@%s:%d/?replicaSet=%s&directConnection=%t"
)

// Mongo is a struct for storing configuration information for connecting to a MongoDB database.
type Mongo struct {
	// User is the username for authenticating with the MongoDB server.
	User string `default:""`
	// Password is the password for authenticating with the MongoDB server.
	Password string `default:""`
	// Host is the hostname or IP address of the MongoDB server.
	Host string `default:"127.0.0.1"`
	// Port is the port number of the MongoDB server.
	Port uint `default:"27017"`
	// DB is the name of the database to connect to.
	DB string `default:"DEFAULT_DB"`
	// DirectConnection specifies whether to connect directly to the server or through a mongos proxy.
	DirectConnection bool `default:"true"`
	// UpgradePath is the path to the database's upgrade script.
	UpgradePath string `default:"migrations"`
	// RSName is the replicaSet name.
	RSName string `default:"rs0"`
}

// GetUpgradePath returns the path to the database's upgrade script.
func (dc Mongo) GetUpgradePath() string {
	return dc.UpgradePath
}

// ConnURL returns the connection URL for the MongoDB server.
func (dc Mongo) ConnURL() string {
	if dc.Password != "" {
		return fmt.Sprintf(fullDBConnFmt, dc.User, dc.Password, dc.Host, dc.Port, dc.RSName, dc.DirectConnection)
	}
	return fmt.Sprintf(simpleDBConnFmt, dc.Host, dc.Port, dc.RSName, dc.DirectConnection)
}

func (dc Mongo) ConnURLWithPrefix() string {
	return fmt.Sprintf("mongodb://%s", dc.ConnURL())
}

// DBName returns the name of the database.
func (dc Mongo) DBName() string {
	return dc.DB
}

// NewMongo returns a new Mongo struct with configuration information read from the environment variables.
// If the DB field is set to "DEFAULT_DB", it will be replaced with the defaultDB parameter.
func NewMongo(defaultDB string) *Mongo {
	var dc Mongo
	if err := envconfig.Process("MONGODB", &dc); err != nil {
		panic(err)
	}
	if dc.DB == "DEFAULT_DB" {
		dc.DB = defaultDB
	}
	return &dc
}

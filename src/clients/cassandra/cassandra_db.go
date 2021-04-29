package cassandra

import (
	"github.com/gocql/gocql"
)

const (
	host        = "127.0.0.1"
	keyspace    = "oauth"
	consistency = gocql.Quorum
)

var (
	clauster *gocql.ClusterConfig
)

func init() {
	//Connect to Cassandra clauster
	clauster = gocql.NewCluster(host)
	clauster.Keyspace = keyspace
	clauster.Consistency = consistency
}

func GetSession() (*gocql.Session, error) {
	return clauster.CreateSession()
}

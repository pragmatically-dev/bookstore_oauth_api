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
	session *gocql.Session
)

func init() {
	//Connect to Cassandra clauster
	clauster := gocql.NewCluster(host)
	clauster.Keyspace = keyspace
	clauster.Consistency = consistency
	var err error
	session, err = clauster.CreateSession()
	if err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}

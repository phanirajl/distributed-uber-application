package cassandra_server

import (
	"github.com/gocql/gocql"
)

func main() {

	// connect to cluster
	cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.3")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum

	session, _ := cluster.CreateSession()

	defer session.Close() // close session

}

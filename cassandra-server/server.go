package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gocql/gocql"
	"github.com/gubrul/distributed-uber-application/kafka-server/server"
	"log"
	"os"
	"os/signal"
)

var (
	CassandraKeySpaces = []string{"users", "transport"}
	CassandraAddr      = []string{"127.0.0.1:9042"}
)

var cassandraCluster = gocql.NewCluster(arrStringToString(CassandraAddr))
var session, sessionErr = cassandraCluster.CreateSession()

func main() {

	//kafkaBrokers := []string{"localhost:9092"}

	cassandraCluster.Keyspace = "users"
	cassandraCluster.Consistency = gocql.Quorum

	if sessionErr != nil {
		log.Println(sessionErr)
		return // exit program if cannot create cassandra session
	}

	defer func() {
		fmt.Println("Cassandra session closed")
		session.Close() // close session
	}()

	// create kafka consumer to topics
	// kafkaConfig := cluster.NewConfig() // new sarama cluster configuration
	// kafkaConfig.Consumer.Return.Errors = true
	// kafkaConfig.Group.Return.Notifications = true

	fmt.Printf("Topics: %+v\n", server.GetTopics()) // print broker's topic

	go func() { // riders topic kafka consumer - process to Cassandra
		msg := server.ConsumeMessage("riders", 0, processRiders)
		log.Printf("Received new driver: %+v\n", msg)
	}()

	go func() { // drivers topic kafka consumer
		msg := server.ConsumeMessage("drivers", 0, processDrivers)
		log.Printf("Received new driver: %+v\n", msg)
	}()

	go func() { // routes topic kafka consumer
		msg := server.ConsumeMessage("routes", 0, processRoutes)
		log.Printf("Received new route: %+v\n", msg)
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt) // os key signal will kill loop

	for {
		select {
		case <-signals:
			{
				return // exit all go routines
			}
		}
	}

}

// consumer routes created in kafaka routes topic
func processRoutes(msg *sarama.ConsumerMessage) {
	fmt.Println(string(msg.Value))
	// insert route into routes table

}

// receive rider ID and stores in Riders cassandra table
func processRiders(msg *sarama.ConsumerMessage) {

	// insert rider to riders table
	// INSERT INTO riders (id,event_time,location_country) VALUES (now(),'2018-10-02 13:05:23.332','Israel');
	query := fmt.Sprintf(`INSERT INTO Riders (id,event_time,location_country) VALUES (%s,'2018-10-02 13:05:23.332','Israel');`, string(msg.Value))

	if err := session.Query(`INSERT INTO users.Riders (id,event_time,location_country) VALUES (?,?,?)`, gocql.TimeUUID(), "2018-10-02 13:05:23.332", "Israel").Exec(); err != nil {
		log.Printf("Error at executing query %+v", query)
	}

}

// receive driver ID and stores in Drivers table
func processDrivers(msg *sarama.ConsumerMessage) {
	// insert driver to drivers table
	fmt.Println(string(msg.Value))
	// INSERT INTO Drivers (id,event_time,location_country,city) VALUES (now(),'2018-10-02 13:05:23.332','Israel','Tel-Aviv');
}

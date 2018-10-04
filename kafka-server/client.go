package main

import (
	"fmt"
	"github.com/uber-go/kafka-client"
	"github.com/uber-go/kafka-client/kafka"
	"github.com/uber-go/tally" // Go metrics interface with fast buffered metrics
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
)

const (
	consumerGoRoutines = 100 // go routines to process messages
)

var (
	client_brokers = map[string][]string{
		"location_cluster":     {"127.0.0.1:9092"}, // default kafka-start port 9092
		"location_dql_cluster": {"127.0.0.1:9092"},
	}

	topicClusterAssignment = map[string][]string{
		"location": {"location_cluster"},
	}
)

func printKafkaMessage(msg kafka.Message) {
	fmt.Printf("Topic: %+v\n", msg.Topic())
	fmt.Printf("Key: %+v\n", msg.Key())
	fmt.Printf("Value: %+v\n", msg.Value())
	fmt.Printf("Time: %+v\n", msg.Timestamp())

}

func processKafkaMessage(msg kafka.Message) error {
	printKafkaMessage(msg)
	var err error // nil
	return err
}

func client() {

	// start kafka-start and zookeeper on your machine

	// First create kafka-start client
	// maps topics to clusters
	client := kafkaclient.New(kafka.NewStaticNameResolver(topicClusterAssignment, client_brokers), zap.NewNop(), tally.NoopScope)

	// Next, setup the consumer config for consuming from a set of topics
	consumerConfig := &kafka.ConsumerConfig{
		TopicList: kafka.ConsumerTopicList{
			kafka.ConsumerTopic{ // Consumer Topic is a combination of topic + dead-letter-queue
				Topic: kafka.Topic{ // Each topic is a tuple of (name, clusterName)
					Name:    "location",
					Cluster: "location_cluster",
				},
				DLQ: kafka.Topic{
					Name:    "location_consumer_dlq",
					Cluster: "location_dlq_cluster",
				},
			},
		},
		GroupName:   "location_consumer",
		Concurrency: consumerGoRoutines, // number of go routines processing messages in parallel
	}
	consumerConfig.Offsets.Initial.Offset = kafka.OffsetNewest // -1
	consumer, err := client.NewConsumer(consumerConfig)        // create new consumer

	if err != nil {
		log.Panic(err)
	}

	// Start consuming
	if err := consumer.Start(); err != nil {
		log.Panic(err)
	}

	sigCh := make(chan os.Signal, 1) // single signal in channel
	signal.Notify(sigCh, os.Interrupt)

	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok != true {
				return // channel closed
			}
			err := processKafkaMessage(msg)
			if err != nil {
				// decline message
				msg.Nack()
			} else {
				// accept message
				msg.Ack()
			}
		case <-sigCh:
			consumer.Stop()
			<-consumer.Closed()
		}
	}

}

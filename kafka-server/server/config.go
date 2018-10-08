package server

import "github.com/Shopify/sarama"

var ( // kafka configurations
	KafkaTopics  = map[string]int{"location": 0, "drivers": 1, "riders": 2, "routes": 3} // kafka topics
	KafkaBrokers = []string{"127.0.0.1:9092"}
	KafkaConfig  = sarama.NewConfig() // default kafka configuration

	CassandraKeySpace = []string{"users"}
	ConsumerGroupId   = []string{"my-consumer"}
)

package server

import "errors"

var (
	ErrorKafkaTopicExists = errors.New("kafka topic exists")
)

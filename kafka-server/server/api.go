package server

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"strings"
)

// Subscribe to topic in partition i
// Receives a topic, partition index and a message processing Function
// Default partition is 0
// runs as a loop that listens to messages
// return number of consumed messages
func ConsumeMessage(topic string, partition int32, processMessage func(msg *sarama.ConsumerMessage)) int { //
	var err error
	// check if topic exists
	if _, exists := KafkaTopics[topic]; exists != true {
		// exist
		log.Printf("Not topic %s exists", topic)
		return 0
	}
	// creates new kafka client
	// creates new async producer to allow produce messages through client
	consumer, err := sarama.NewConsumer(KafkaBrokers, KafkaConfig)

	if err != nil {
		log.Println(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)

	if err != nil {
		log.Println(err) // could not consume partition in topic
	}

	defer func() { // kill consumer client
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt) // os key signal will kill loop

	consumed := 0 // count how many messages consumed

ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			// process consumed message
			fmt.Printf("msg: %s", msg.Value)

			processMessage(msg) // process message RUN AS GO ROUTINE?? KILL ROUTINE WHEN DONE EXECUTION

			log.Printf("Consumed message offset %d\n", msg.Offset)
			// mark message as processed
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}
	log.Printf("[%s]Consumed: %d\n", topic, consumed)
	return consumed
}

// produce single message to topic
// creates AsyncProducer on top of sarama kafka client
func ProduceMessage(topic string, message string) {

	// creates new kafka client
	// creates new async producer to allow produce messages thrgh client
	producer, err := sarama.NewAsyncProducer(KafkaBrokers, KafkaConfig)

	if err != nil {
		log.Println(err)
	}

	defer func() { // kill producer client
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	//	producer.Input() <- &sarama.ProducerMessage{ // produce message
	//		Topic: topic,
	//		Key: sarama.StringEncoder("The key of the message"), // encoded to byte array
	//		Value: sarama.StringEncoder(message), // encoded to byte array
	//	}
	//	log.Printf("[%s] Sent message: %s",topic,message) // log message

	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(message)}

}

// streams messages to topic from a single AsyncProducer on top of sarama kafka client
// stops on OS signal interruption
// runs as a loop and not as go routine?
func ProducerStreamMessage(brokers []string, config *sarama.Config, topic string, message string) (int, int) {

	// create new kafka client
	// creates new async producer to allow produc messages thrgh client
	producer, err := sarama.NewAsyncProducer(brokers, config)

	if err != nil {
		log.Println(err)
	}

	defer func() { // kill producer on scope exit
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown for loop.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var enqueued, countErrors int

ProducerLoop:
	for {
		select {
		case producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(message)}:
			enqueued++
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
			countErrors++
		case <-signals:
			break ProducerLoop
		}
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, countErrors)

	return enqueued, countErrors
}

// print all topics of kafka brokers
func GetTopics() string {
	client, err := sarama.NewClient(KafkaBrokers, KafkaConfig)

	if err != nil {
		log.Println("Failed to create client")
		return ""
	}

	defer func() { // close client
		client.Close()
	}()

	topics, _ := client.Topics()
	return strings.Join(topics, ",") // return topics

}

// create kafka topic
// returns an error of topic exists
func CreateTopic(topic string) error {
	var err error

	return err
}

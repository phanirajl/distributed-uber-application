from kafka import SimpleProducer, KafkaClient

# connect to kafka
kafka = KafkaClient("localhost:9092")

producer = SimpleProducer(kafka)

# assign a topic
topic = "getTaxi"

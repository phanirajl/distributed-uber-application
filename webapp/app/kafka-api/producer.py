from kafka import SimpleProducer, KafkaClient, KafkaProducer

# connect to kafka-api-start
kafka = KafkaClient("localhost:9092")

producer = SimpleProducer(kafka)

# assign a topic
topic = "getTaxi"

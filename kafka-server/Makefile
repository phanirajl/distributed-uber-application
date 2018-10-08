SCALA_VERSION?= 2.11
KAFKA_VERSION?= 2.0.0
KAFKA_DIR= kafka_$(SCALA_VERSION)-$(KAFKA_VERSION)
PKG=$(shell go list ./... | grep -v vendor)

default:
	go vet $(PKG)

test:
	go test

vet:
	go vet $(PKG)

build:
	go build ./...
	./kafka-server

run kafka:
	./kafka-server

doc:
	echo "Kafka documentation"
	$(PKG)/README.md

server: # triger kafka server
	go build -o bin/kafka-server
	bin/kafka-server --brokers 0,1,2

client: # trigger kafka client
	go build -o bin/kafka-server
	bin/kafka-server --brokers 0,1,2

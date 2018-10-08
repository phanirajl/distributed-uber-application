SCALA_VERSION?= 2.11
CASSANDRA_VERSION?= 3.11.3
CASSANDRA_DIR= apache-cassandra-$(CASSANDRA_VERSION)
PKG=$(shell go list ./... | grep -v vendor)

default:
	go vet $(PKG)

test:
	go test

vet:
	go vet $(PKG)

build:
	go build -o ./cassandra-server
	./cassandra-server

run cassandra:
	./cassandra-server

doc:
	echo "Cassandra documentation"
	$(PKG)/README.md


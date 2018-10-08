
#### Install and Configure Cassandra

Cassandra's main goal is to handle big data workloads across multiple nodes without any 
single point of failure. Cassandra has p2p distributed system across it's nodes.

Similar to Kafka Cassandra support Data Replication paradigm.


##### Quick overview about Cassandra components:
- __Node__ - data stored in.
- __Data Center__ - collection of nodes.
- __Cluster__ - componenet that contains one or more data centers.
- __Commit log__ - crash recovery mechanism in Cassandra.


##### Keyspace 
Keyspace is the outermost container of data in Cassandra.
- __Replication Factor__ - num of machines in cluster replicated.
- __Replica placement strategy__ - strategy to place replicas in the ring.
- __Column families__ - container of a collection of rows.

Cassandra deals with unstructured data. 
Row is a unit of replication in Cassandra. 
Column is a unit of storage in Cassandra.
Relationships are represented using collections.

##### Download and install Cassandra
```bash
$ wget http://apache.mivzakim.net/cassandra/3.11.3/apache-cassandra-3.11.3-bin.tar.gz
$ tar zxvf apache-cassandra-3.11.3-bin.tar.gz
$ mkdir Cassandra
$ mv apache-cassandra-3.11.3 cassandra.
```

###### Make sure you are running Java SE JDK 8
Later jdk version are incompatible with cassandra failing the JVM to run.
Comment the following lines in conf/jvm.options file:
-XX:ThreadPriorityPolicy=42
-XX:+UseParNewGC 
Run this in command line:
	JAVA_MAJOR_VERSION=$($JAVA -version 2>&1 | sed -E -n 's/.* version "([^.-]*).*"/\1/p')

##### Start Cassandra 
```bash
$ bin/cassandra
```
Open another tab to start Cassandra CLI:
```bash
$ bin/cqlsh
```

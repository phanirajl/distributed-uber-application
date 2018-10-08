
#### Configure Kafka and Zookeeper
__Install kafka__ (http://apache.mivzakim.net/kafka/2.0.0/kafka_2.11-2.0.0.tgz)

__Install zookeeper__ (http://apache.mivzakim.net/zookeeper/)

__Run zookeeper:__ 
```bash
 $ cd zookeeper-3.4.10 
 $ bin/zkServer.sh start 
 $ bin/zkCli.sh
```
 

__Run kafka:__ 
```bash
 $ cd kafka_2.11-2.0.0 
 $ bin/kafka-server-start.sh config/server.properties 
```
  
__Type jps in kafka terminal:__ 
```bash
 $ jps 
   14644 Jps 
   10006 QuorumPeerMain # is zookeeper deamon 
   14377 Kafka 
   461 
```
   
Path with brew: <br />
Kafka path : /usr/local/Cellar/kafka/2.0.0/ <br />
Zookeeper path: /usr/local/Cellar/zookeeper/3.4.12/


##### Lets set topics for our broker configurations 
__Run in kafka directory:__ 
```bash
$ bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 
         --partitions 1 --topic location 
```
      
__To get list of all topics:__ 
```bash
 $ bin/kafka-topics.sh --list --zookeeper localhost:2181
```

__Produce our rider's location:__ 
```bash
$ bin/kafka-console-producer.sh --broker-list localhost:9092 --topic location 
       $ {latitude:40.124124124, longitude:123.23123123213}
```
    
__Create a consumer (driver) to receive rider's location:__ 
```bash
 $ bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic location --from-beginning
       __Output:__ 
       {latitude:40.124124124, longitude:123.23123123213}
```
      

##### Now let's expand our broker cluster to 3 nodes:

__Create some new server-1 and server-2 configurations:__
``` bash 
   $ cp config/server.properties config/server-1.properties
   $ cp config/server.properties config/server-2.properties
   $ vim config/server-1.properties
      broker.id=1
      listeners=PLAINTEXT://:9093
      log.dirs=/tmp/kafka-logs-1
   $ vim config/server-2.properties
      broker.id=2
      listeners=PLAINTEXT://:9094
      log.dirs=/tmp/kafka-logs-2
```
__So in total we have 3 nodes listening to ports 9092,9093,9094.__

___Let's start these nodes:___
```bash
   $  bin/kafka-server-start.sh config/server-1.properties &
   $  bin/kafka-server-start.sh config/server-2.properties &
```
__Our old topic location has only replication factor of 1 but because we added 2 more brokers we need to upgrade the replication
   factor to three. So let's delete and create a new one with 3 replicas__
```bash 
   $  bin/kafka-topics.sh --delete --zookeeper localhost:2181 --topic location
   $  bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 3 --partitions 1 --topic location  
```

__Gets topic's brokers information:__
```bash
  $  bin/kafka-topics.sh --describe --zookeeper localhost:2181 --topic location
     Topic:location	PartitionCount:1	ReplicationFactor:3	Configs:
      Topic: location	Partition: 0	Leader: 2	Replicas: 2,0,1	Isr: 2,0,1
```
   
  First line tells us how many partitions we have in this case 1 and how many replicas (3).<br />
  Seconds line tells us information about that single partition - ran by 3 node brokers 0,1,2<br />
  * "leader" is the node responsible for all reads and writes for the given partition. Each node will be the leader for a randomly selected portion of the partitions.<br />
  * "replicas" is partitions list of node brokers<br />

__Publish few messages to our location topic:__<br />
```bash
 $ bin/kafka-console-producer.sh --broker-list localhost:9092 --topic location
        {long:38.895100,lat:-77.036400,city:"Washington DC"}
        {long:34.781769,lat:32.085300,city:"Tel Aviv"}
        {long:34.989571,lat:32.794044,city:"Haifa"}
```
__Consume our messages from our location topic:__<br />
```bash
$ bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --from-beginning --topic location
        {long:38.895100,lat:-77.036400,city:"Washington DC"}
        {long:34.781769,lat:32.085300,city:"Tel Aviv"}
        {long:34.989571,lat:32.794044,city:"Haifa"}
```

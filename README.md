Distributed uber alike application
-----------------------------------
[![License](https://img.shields.io/badge/license-Apache-blue.svg)](https://github.com/gubrul/distributed-uber-application/LICENSE)
[![PyPI version](https://badge.fury.io/py/catboost.svg)](https://badge.fury.io/py/catboost)

An uber data pipeline web application. Kafka messaging broker, flask web application serving frontend app, Spark streaming and Cassandra database.

The project is divided into 3 sub components, web application, kafka server and cassandra server. Each component is Dockerized allowing distribute our
backend load which we will discuss later. To manage these docker containers we will kubernetes.




### Application Architecture 

####6 distributed "micro" services: 
   - __kafka server__
        - messaging broker for our riders and drivers communication
   - __cassandra server__
        - store drivers real time location
        - store rider's location
        - store rides information
   - __map-location__
        - process location
   - __map-route__
        - process route from A to B
        - predict route duration
   - __which-cab__
        - analyze ride specification and choose the closest cab to send
   - __web application__
        - display map and connect producers and consumers via web interface
        - display in real time drivers and riders

#### Web application 
##### Requiremenets 
- python 3.7.0 
- pip3

1) Create virtual environment and install pip requirements.
```bash
$distributed-uber-application: cd webapp/app
$distributed-uber-application: virtualenv env
$distributed-uber-application: source env/bin/activate
$distributed-uber-application: pip install -r requirements.txt
$distributed-uber-application: python manage.py
```


#### Kafka overview 
##### Topics:
   - __location__ {long:34.781769, lat:32.085300, city:"Tel Aviv"}
   - __ride__ {duration:10min, traffic_density:3,
   - __passenger__ {numOfPassengers:1,disable:false, numOfLaggage:2}

__Traffic Density Parameter (TDP):__ <br />
   - Calculate average speed a ride takes in certain road section and increase 1 density weight per 5km/h the car move slower.
   - Goes from 0 to 10 (0-free, 10-traffic jam}


#### Configure Kafka and Zookeeper
__Install kafka__ (http://apache.mivzakim.net/kafka/2.0.0/kafka_2.11-2.0.0.tgz)

__Install zookeeper__ (http://apache.mivzakim.net/zookeeper/)

__Run zookeeper:__ <br />
```bash
 $ cd zookeeper-3.4.10 
 $ bin/zkServer.sh start 
 $ bin/zkCli.sh
```
 

__Run kafka:__ <br />
```bash
 $ cd kafka_2.11-2.0.0 
 $ bin/kafka-server-start.sh config/server.properties 
```
  
__Type jps in kafka terminal:__ <br />
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
$ bin/kafka-console-producer.sh --broker-list localhost:9092 --topic location <br />
       $ {latitude:40.124124124, longitude:123.23123123213}
```
    
__Create a consumer (driver) to receive rider's location:__ 
```bash
 $ bin/kafka-console-consumer.sh --bootstrap-server localhost:2181 —topic location <br />
         --from-beginning <br />
       __Output:__ <br />
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
__So in total we have 3 nodes listening to ports 9092,9093,9094.__<br />

___Let's start these nodes:___
```bash
   $  bin/kafka-server-start.sh config/server-1.properties &
   $  bin/kafka-server-start.sh config/server-2.properties &
```
__Our old topic location has only replication factor of 1 but because we added 2 more brokers we need to upgrade the replication
   factor to three. So let's delete and create a new one with 3 replicas__
```bash 
   $  bin/kafka-topics.sh --delete --zookeeper localhost:2181 --topic location
   $  bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 3 --partitions 1 --topic location<br />   
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

### Run Cassandra 
More details soon...
### Deployment Instructions 

We will need to spin all our docker containers within kubernetes.






Distributed uber alike application
-----------------------------------

[![License: GNU](https://img.shields.io/badge/license-GNU-blue.svg)](https://github.com/gubrul/distributed-uber-application/LICENSE)
[![Kafka 2.0.0](https://img.shields.io/badge/kafka-2.0.0-brightgreen.svg)](http://apache.mivzakim.net/kafka/2.0.0/kafka-2.0.0-src.tgz)
[![Cassandra 3.11.3](https://img.shields.io/badge/kafka-3.11.3-brightgreen.svg)](http://apache.mivzakim.net/cassandra/3.11.3/apache-cassandra-3.11.3-bin.tar.gz)
[![Java JDK 8](https://img.shields.io/badge/JDK-6%20%7C%207%20%7C%208-red.svg)](https://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html)
[![Python 3.5-3.7](https://img.shields.io/badge/python-3.5%20%7C%203.6%20%7C%203.7-blue.svg)]()
[![Golang 1.11](https://img.shields.io/badge/Golang-1.11-blue.svg)]()<br/>
An uber data pipeline web application. Kafka messaging broker, flask web application serving frontend app, Spark streaming and Cassandra database.

The project is divided into 3 sub components, web application, kafka server and cassandra server. Each component is Dockerized allowing distribute our
backend load which we will discuss later. To manage these docker containers we will kubernetes.




### Application Architecture 

#### 6 distributed "micro" services: 
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
   - __weather__
        -  Get weather in route region    
   - __web application__
        - display map and connect producers and consumers via web interface
        - display in real time drivers and riders

Visual architecture:
![alt text](https://github.com/gubrul/distributed-uber-application/blob/master/docs/architecture.png)

### [Kafka Installation and Configuration Docs](https://github.com/gubrul/distributed-uber-application/blob/master/docs/KAFKA.md)
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

__Traffic Density Parameter (TDP):__ 
   - Calculate average speed a ride takes in certain road section and increase 1 density weight per 5km/h the car move slower.
   - Goes from 0 to 10 (0-free, 10-traffic jam}


### [Cassandra Instalation/Configuration and overview](https://github.com/gubrul/distributed-uber-application/blob/master/docs/CASSANDRA.md)

### Data Modeling

Let's create a Keyspace:
```bash
cqlsh> CREATE KEYSPACE Users
   ... WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 3};
cqlsh> DESCRIBE keyspaces
       
       users  system_auth  system_distributed
       system_schema   system       system_traces  
```
Our replication factor is 3 which means we will have 3 nodes in our cluster in a Simple Strategy node 
ring position.

First let's access our keyspace and create riders table

```sql
cqls> USE users;
cqlsh:users> CREATE TABLE Riders(
             id uuid,
             event_time timestamp,
             location_country text,
             city text,
             latitude_e6 bigint,
             longitude_e6 bigint,
             disabled boolean,
             PRIMARY KEY (id)
             );
```
insert into stuff (uid, name) values(now(), 'my name');
Let's insert some data:
```sql
cqlsh:users> INSERT INTO Riders (id,event_time,location_country) VALUES (now(),'2018-10-02 13:05:23.332','Israel');
cqlsh:users> INSERT INTO Riders (id,event_time,location_country) VALUES (now(),'2018-10-02 13:05:23.332','Usa');
cqlsh:users> INSERT INTO Riders (id,event_time,location_country) VALUES (now(),'2018-10-02 13:05:23.332','Germany');
cqlsh:users> INSERT INTO Riders (id,event_time,location_country) VALUES (now(),'2018-10-02 13:05:23.332','Belgium');
```
Select all our tables data:
```sql
cqlsh:users> SELECT * FROM Riders;

 id                                   | city | disabled | event_time                      | latitude_e6 | location_country | longitude_e6
--------------------------------------+------+----------+---------------------------------+-------------+------------------+--------------
 216faa10-ca0a-11e8-a83d-ebd8dd36fbd2 | null |     null | 2018-10-02 10:05:23.332000+0000 |        null |          Belgium |         null
 1e041d70-ca0a-11e8-a83d-ebd8dd36fbd2 | null |     null | 2018-10-02 10:05:23.332000+0000 |        null |          Germany |         null
 1450e650-ca0a-11e8-a83d-ebd8dd36fbd2 | null |     null | 2018-10-02 10:05:23.332000+0000 |        null |           Israel |         null
 19c5d780-ca0a-11e8-a83d-ebd8dd36fbd2 | null |     null | 2018-10-02 10:05:23.332000+0000 |        null |              Usa |         null

(4 rows)
```
We will need to create a table for our drivers as well:
```sql
cqls> USE users;
cqlsh:users> CREATE TABLE Drivers(
             id uuid,
             event_time timestamp,
             location_country text,
             city text,
             disable_compatible boolean,
             PRIMARY KEY (id)
             );
```
And again insert some data to the table and print it:
```sql
cqlsh:users> INSERT INTO Drivers (id,event_time,location_country,city) VALUES (now(),'2018-10-02 13:05:23.332','Israel','Tel-Aviv');
cqlsh:users> INSERT INTO Drivers (id,event_time,location_country,city) VALUES (now(),'2018-10-02 13:05:23.332','Usa','Washington');
cqlsh:users> INSERT INTO Drivers (id,event_time,location_country,city) VALUES (now(),'2018-10-02 13:05:23.332','Germany','Berlin');
cqlsh:users> INSERT INTO Drivers (id,event_time,location_country,city) VALUES (now(),'2018-10-02 13:05:23.332','Belgium','Brussels');
cqlsh:users> SELECT * FROM Drivers;

 id                                   | city       | disable_compatible | event_time                      | location_country
--------------------------------------+------------+--------------------+---------------------------------+------------------
 31fc9d10-ca0b-11e8-a83d-ebd8dd36fbd2 | Washington |               null | 2018-10-02 10:05:23.332000+0000 |              Usa
 31faef60-ca0b-11e8-a83d-ebd8dd36fbd2 |   Tel-Aviv |               null | 2018-10-02 10:05:23.332000+0000 |           Israel
 3289c1e0-ca0b-11e8-a83d-ebd8dd36fbd2 |   Brussels |               null | 2018-10-02 10:05:23.332000+0000 |          Belgium
 31fdae80-ca0b-11e8-a83d-ebd8dd36fbd2 |     Berlin |               null | 2018-10-02 10:05:23.332000+0000 |          Germany

(4 rows)
```


Let's create another keyspace for all of our transportation data:
```sql
cqlsh> CREATE KEYSPACE Transport
   ... WITH replication = {'class':'SimpleStrategy','replication_factor':3};
cqlsh> DESCRIBE Transport
CREATE KEYSPACE transport WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '3'}  
AND durable_writes = true;
```


### Deployment Instructions 
We will need to spin all our docker containers within kubernetes.






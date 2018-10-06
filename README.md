Distributed uber alike application
-----------------------------------

[![License: GNU](https://img.shields.io/badge/license-GNU-blue.svg)](https://github.com/gubrul/distributed-uber-application/LICENSE)
[![Kafka](https://img.shields.io/badge/kafka-2.0.0-brightgreen.svg)]()
[![Python 3.5-3.7](https://img.shields.io/badge/python-3.5%20%7C%203.6%20%7C%203.7-blue.svg)]()
[![Golang 1.11](https://img.shields.io/badge/Golang-1.11-blue.svg)]()
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


### Run Cassandra 
More details soon...
### Deployment Instructions 

We will need to spin all our docker containers within kubernetes.






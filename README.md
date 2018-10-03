Distributed Uber alike application
-----------------------------------

An uber data pipeline web application. Kafka messaging broker, flask web application serving frontend app, Spark streaming and Cassandra database.

The project is divided into 3 sub components, web application, kafka server and cassandra server. Each component is Dockerized allowing distribute our
backend load which we will discuss later. To manage these docker containers we will kubernetes.


Instructions:

We will need to spin all our docker containers within kubernetes.


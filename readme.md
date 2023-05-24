# Learn Kafka Golang

## Prerequisites
- Golang 1.20.4
- Kafka 3.0 https://kafka.apache.org/
- zookeper

## Source
- https://www.sohamkamani.com/install-and-run-kafka-locally/
- https://www.sohamkamani.com/golang/working-with-kafka/

## Step
- make sure in root folder where kafka downloaded

- run zookeper
  ```
  bin/zookeeper-server-start.sh config/zookeeper.properties
  ```

- run kafka
  ```
  bin/kafka-server-start.sh config/server.properties
  ```

- However, there are three properties that have to be unique for each broker instance:
  ```
  File: kafka_2.13-3.0.0/config/server.properties

  # The id of the broker. This must be set to a unique integer for each broker.
  broker.id=1

  # The address the socket server listens on. It will get the value returned from
  listeners=PLAINTEXT://:9092

  # A comma separated list of directories under which to store log files
  log.dirs=/tmp/kafka-logs
  ```

- creating topics
  ```
  bin/kafka-topics.sh --create --topic my-kafka-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
  ```
  - `partitions` lets you decide how many brokers you want your data to be split between
  - `replication-factor` describes how many copies of you data you want (in case one of the brokers goes down, you still have your data on the others). Since we set this value to 2, our data will have two copies of itself on any two of the brokers
  - `bootstrap-server` points to the address of any one of our active Kafka brokers. Since all brokers know about each other through Zookeeper, it doesn’t matter which one you choose.

- listing topics
  ```
  bin/kafka-topics.sh --list --bootstrap-server localhost:9092
  ```

- To get more details on a topic, you can use the --describe argument with the same command:
  ```
  bin/kafka-topics.sh --describe --topic my-kafka-topic --bootstrap-server localhost:9092
  ```

- To start the console producer, run the command:
  ```
  bin/kafka-console-producer.sh --broker-list localhost:9092 --topic my-kafka-topic
  ```
  - `broker-list` points the producer to the addresses of the brokers that we just provisioned
  - `topic` specifies the topic you want the data to go to.

- To start a consumer, run the command:
  ```
  bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic my-kafka-topic --from-beginning
  ```
  - `bootstrap-serve`r` can be any one of the brokers in the cluster
  - `topic` should be the same as the topic under which you producers inserted data into the cluster.
  - `from-beginning` tells the cluster that you want all the messages that it currently has with it, even messages that we put into it previously.
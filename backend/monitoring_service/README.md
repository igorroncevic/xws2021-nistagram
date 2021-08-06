# Notes

## Container docs

- Kafka: https://hub.docker.com/r/wurstmeister/kafka

## Basic commands

### Entering Kafka container

- ```docker exec -it kafka /bin/sh```

## Topics

### Creating a topic

- ```cd /opt/kafka_2.13-2.7.0```
- ```./bin/kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic topicname```

### Listing topics

- ```./bin/kafka-topics.sh --list --zookeeper zookeeper:2181```
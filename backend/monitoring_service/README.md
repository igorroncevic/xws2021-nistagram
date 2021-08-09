# Notes

## Quick Start guide
For Ništagram, you will need to create two topics: "user-events" and "performance" using commands below.

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

### Test read/write to a topic
- Write: ```bin/kafka-console-producer.sh --topic quickstart-events --bootstrap-server localhost:9092```
- Read: ```bin/kafka-console-consumer.sh --topic quickstart-events --from-beginning --bootstrap-server localhost:9092```
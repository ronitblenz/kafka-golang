# Apache Kafka Implementation using GoLang

This project demonstrates how to set up Apache Kafka using Docker and implement a producer and consumer service in Golang.

## Folder Structure

```
kafka-golang/
├── docker-compose.yml
├── producer/
│   ├── go.mod
│   └── main.go
└── consumer/
    ├── go.mod
    └── main.go
```

## Prerequisites

- Docker and Docker Compose installed
- Golang installed

## Setup and Run

1. **Clone the repository:**

   ```sh
   git clone https://github.com/ronitblenz/kafka-golang.git
   cd kafka-golang
   ```

2. **Start Kafka and Zookeeper using Docker Compose:**

   ```sh
   docker-compose up
   ```

3. **Set up and run the Producer:**

   ```sh
   cd producer
   go mod init producer
   go mod tidy
   go run main.go
   ```

4. **Set up and run the Consumer:**

   Open a new terminal and execute:

   ```sh
   cd kafka-golang-example/consumer
   go mod init consumer
   go mod tidy
   go run main.go
   ```

The producer will send messages to the Kafka topic, and the consumer will receive and print those messages.

## Sample Output (Consumer Side)

```sh
2024/08/06 14:02:26 0 msg from producer
2024/08/06 14:02:26 1 msg from producer
2024/08/06 14:02:26 2 msg from producer
2024/08/06 14:02:26 3 msg from producer
2024/08/06 14:02:26 4 msg from producer
```

### Access Kafka Broker

You can use Kafka's command-line tools to interact with the Kafka broker.

1. **Run Kafka container's bash:**

   ```sh
   docker exec -it broker /bin/bash
   ```

2. **List topics:**

   Once inside the container, use the following command to list all topics:

   ```sh
   kafka-topics --list --bootstrap-server localhost:9092
   ```

3. **Describe a topic:**

   To see details about a specific topic:

   ```sh
   kafka-topics --describe --topic msg --bootstrap-server localhost:9092
   ```

4. **Consume messages:**

   You can also consume messages from a topic using the command line to see the real-time data:

   ```sh
   kafka-console-consumer --bootstrap-server localhost:9092 --topic msg --from-beginning
   ```

5. **Produce messages:**

   To produce messages to a topic:

   ```sh
   kafka-console-producer --bootstrap-server localhost:9092 --topic msg
   ```
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
    topic := "msg"
    // setting the consumer
    consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "go",
        "auto.offset.reset": "smallest"})
    if err != nil {
        log.Println(err)
    }
    // subscribing to the same topic as producer
    err = consumer.Subscribe(topic, nil)
    if err != nil {
        log.Println(err)
    }
    for {
        ev := consumer.Poll(100)
        switch e := ev.(type) {
        case *kafka.Message:
            log.Println(string(e.Value))
       // consuming message
        case kafka.Error:
            fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
        }
    }
}

package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
   // creating producer
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "client.id":         "cli",
        "acks":              "all"})

    if err != nil {
        fmt.Printf("Failed to create producer: %s\n", err)
        os.Exit(1)
    }
// defining the name of the topic
    topic := "msg"

    delivery_chan := make(chan kafka.Event, 10000)

    for i := 0; i < 5; i++ {

        value := fmt.Sprintf("%d Hello InfraCloud!", i)
// writing message to a topic
        err := p.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Value:          []byte(value)},
            delivery_chan,
        )
        if err != nil {
            log.Println(err)
        }
// this will block the execution until producing message is done.
        <-delivery_chan
 // to show that how it would perform if it's a time taking event.
        time.Sleep(time.Second * 3)

    }
}

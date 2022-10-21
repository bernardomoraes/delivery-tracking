package main

import (
	"fmt"
	"log"

	kafkaApp "github.com/bernardomoraes/delivery-tracking/application/kafka"
	"github.com/bernardomoraes/delivery-tracking/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("Ola mundo", os.Getenv("kafkaProduceTopic"), producer)

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafkaApp.Produce(msg)
	}
}

package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{MsgChan: msgChan}
}

func (k *KafkaConsumer) Consume() {
	// basic consumer config
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	// Creating new consumer instance
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kafka message:" + err.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopic")}
	c.SubscribeTopics(topics, nil) // rebalance nil is = automatic

	log.Println("Kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1) // -1 = infinite await
		if err == nil {
			fmt.Println("Message on Kafka: ", string(msg.Value))
			k.MsgChan <- msg
		}
	}
}

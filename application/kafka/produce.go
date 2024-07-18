package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bernardomoraes/delivery-tracking/application/route"
	"github.com/bernardomoraes/delivery-tracking/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()
	err := json.Unmarshal(msg.Value, &route)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(1 * time.Second)
	}

}

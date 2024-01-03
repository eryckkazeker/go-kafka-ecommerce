package main

import (
	"context"
	"encoding/json"
	"fmt"

	"com.ekazeker/mail-service/models"
	"github.com/segmentio/kafka-go"
)

func main() {

	consumer("golang-test", "mail-service")

}

func consumer(topic string, consumer string) {
	fmt.Println("Starting to consume topic [", topic, "]")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: consumer,
		Topic:   topic,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		var order = models.Order{}
		json.Unmarshal(m.Value, &order)
		// fmt.Printf("[%s]: %s \n", topic, m.Value)
		fmt.Printf("New order received, Id: %s \n", order.Id)
	}
}

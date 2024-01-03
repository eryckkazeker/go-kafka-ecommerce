package main

import (
	"context"
	"encoding/json"
	"fmt"

	"com.ekazeker/warehouse-service/models"
	"github.com/segmentio/kafka-go"
)

func main() {

	consumer("golang-test", "warehouse-service")

}

func consumer(topic string, groupId string) {
	fmt.Println("Starting to consume topic [", topic, "]")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: groupId,
		Topic:   topic,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		var order = models.Order{}
		json.Unmarshal(m.Value, &order)

		for _, v := range order.Products {
			fmt.Printf("Product reservation Id: %s Quantity: %d\n", v.Id, v.Quantity)
		}
	}
}

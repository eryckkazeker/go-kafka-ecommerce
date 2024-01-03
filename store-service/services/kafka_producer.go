package services

import (
	"context"
	"encoding/json"

	"com.ekazeker/store-service/models"
	"github.com/segmentio/kafka-go"
)

func SendOrder(order models.Order) {
	w := kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "golang-test",
		Balancer: &kafka.LeastBytes{},
	}

	ord, _ := json.Marshal(order)

	w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(order.Id),
			Value: ord,
		})

}

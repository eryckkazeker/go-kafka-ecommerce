package models

import (
	"github.com/google/uuid"
)

type Order struct {
	Id          string         `json:"id"`
	Products    []ProductStock `json:"products"`
	TotalAmount int            `json:"total_amount"`
}

func NewOrder() *Order {
	return &Order{
		Id:          uuid.NewString(),
		Products:    []ProductStock{},
		TotalAmount: 0,
	}
}

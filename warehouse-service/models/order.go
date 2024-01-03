package models

type Order struct {
	Id       string    `json:"id"`
	Products []Product `json:"products"`
}

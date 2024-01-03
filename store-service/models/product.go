package models

type Product struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Price       int    `json:"price"` // value stored in cents
}

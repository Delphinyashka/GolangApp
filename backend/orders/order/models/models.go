package models

type Order struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	ClientID  int `json:"client_id"`
	Quantity  int `json:"quantity"`
}

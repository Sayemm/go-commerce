package models

type CartItem struct {
	ID        int `json:"id" db:"id"`
	ProductID int `json:"product_id" db:"product_id"`
	Quantity  int `json:"quantity" db:"quantity"`
}

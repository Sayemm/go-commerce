package models

type Product struct {
	ID       int     `json:"id" db:"id"`
	Name     string  `json:"name" db:"name"`
	Price    float64 `json:"price" db:"price"`
	Quantity int     `json:"quantity" db:"quantity"`
}

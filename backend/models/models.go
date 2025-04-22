package models

type Product struct {
	ID    int     `db:"id" json:"id"`
	Name  string  `db:"name" json:"name"`
	Price float64 `db:"price" json:"price"`
}

type CartItem struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Order struct {
	ID    int        `db:"id" json:"id"`
	Items []CartItem `json:"items"`
	Total float64    `db:"total"`
}

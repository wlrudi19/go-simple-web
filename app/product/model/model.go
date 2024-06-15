package model

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
	Stok        int    `json:"stok"`
}

type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
	Stok        int    `json:"stok"`
}

type ProductRequest struct {
	Id int `json:"id"`
}

type FindProductResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
	Stok        int    `json:"stok"`
}

type UpdateProductRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Amount      *string `json:"amount"`
	Stok        *int    `json:"stok"`
}

type Order struct {
	Id        int    `json:"id,omitempty"`
	ProductID int    `json:"product_id"`
	UserID    int    `json:"user_id"`
	Total     int    `json:"total"`
	Amount    string `json:"amount"`
	Status    string `json:"status"`
	CollectId []int  `json:"collect_id,omitempty"`
}

type OrderSummary struct {
	Data       []Order `json:"data"`
	Kupon      int     `json:"kupon"`
	TotalBayar float64 `json:"bayar"`
}

type BulkUpdateOrder struct {
	ProductID     int    `json:"product_id"`
	Total         int    `json:"total"`
	ProductUpdate bool   `json:"product_update"`
	Status        string `json:"status"`
	CollectId     []int  `json:"collect_id"`
}

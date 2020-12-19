package create

type ReqCreateProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int8    `json:"quantity"`
}
package create

type ReqCreateProduct struct {
	name        string  `json:"name"`
	description string  `json:"description"`
	price       float64 `json:"price"`
	quantity    int8    `json:"quantity"`
}

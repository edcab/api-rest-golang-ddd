package entities

import "fmt"

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int8
}

func NewProduct(name string, description string, price float64, quantity int8) (*Product, error) {

	if name == "" || description == "" || price == 0 || quantity == 0 {
		return nil, fmt.Errorf(
			"Product value are not valid, name: %v, description: %v, price: %v, quantity: %v",
			name,
			description,
			price,
			quantity,
		)
	}

	return &Product{Name: name, Description: description, Price: price, Quantity: quantity}, nil
}




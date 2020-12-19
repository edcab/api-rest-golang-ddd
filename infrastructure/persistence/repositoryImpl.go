package persistence

import (
	"api-rest-golang-ddd/domain/entities"
	"fmt"
)

type CrudProduct interface {
	SaveProduct() error
}

type RepositoryStorageProduct struct {
	product entities.Product
}

func (r RepositoryStorageProduct) SaveProduct() error {
	fmt.Print("Product to save", r.product)
	fmt.Println("Product created successfully")
	return nil
}

func NewCrudProduct(product entities.Product) *RepositoryStorageProduct {
	return &RepositoryStorageProduct{ product: product }
}


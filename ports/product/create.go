package product

import (
	"api-rest-golang-ddd/domain/dto"
	"api-rest-golang-ddd/usecase"
)

type IProductPort interface{
	Create(product dto.Product) error
}

type PortProductImpl struct{

}

func (p PortProductImpl) Create(product dto.Product) error {
	productUseCase := usecase.NewCreateProductUseCase()

	productUseCase.CreateProduct(product)

	return nil
}

func NewCreateProductPort() *PortProductImpl {
	return &PortProductImpl{}
}
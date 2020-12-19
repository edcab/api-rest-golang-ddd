package usecase

import (
	"api-rest-golang-ddd/domain/dto"
	"api-rest-golang-ddd/domain/entities"
	"api-rest-golang-ddd/infrastructure/persistence"
)

type CreateProductUseCase interface {
	CreateProduct(productDto dto.Product)
}

type CreateProductUseCaseImpl struct{

}

func (uc *CreateProductUseCaseImpl) CreateProduct(product dto.Product) error{

	newProduct, err := entities.NewProduct(product.Name, product.Description, product.Price, product.Quantity)

	if err != nil{
		return err
	}

	crudProduct := persistence.NewCrudProduct(*newProduct)

	err = crudProduct.SaveProduct()

	if err != nil {
		return err
	}

	return nil
}

func NewCreateProductUseCase() *CreateProductUseCaseImpl {
	return &CreateProductUseCaseImpl{}
}
package product

import (
	"api-rest-golang-ddd/domain/dto"
	"testing"
)

func Test_CreatePort(t *testing.T) {
	productDto := dto.Product{
		Name:        "ProductTest",
		Description: "Product for testing purpose",
		Price:       10,
		Quantity:    1,
	}

	port := NewCreateProductPort()

	got := port.create(productDto)

	if got != nil{

		t.Errorf("Error testing create port, want nil, got %v", got)

	}

}
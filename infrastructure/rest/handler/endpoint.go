package handler

import (
	"api-rest-golang-ddd/domain/dto"
	"api-rest-golang-ddd/infrastructure/rest/model/create"
	"api-rest-golang-ddd/ports/product"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

func MakeCreateProductEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var body create.ReqCreateProduct
		req := request.(*http.Request)

		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			return nil, err
		}

		port := product.NewCreateProductPort()

		newProduct, err := dto.NewProduct(body.Name, body.Description, body.Price, body.Quantity)

		if err != nil {
			return nil, fmt.Errorf("Error constructing a DTO of Product, error: %v", err.Error())
		}

		err = port.Create(*newProduct)

		if err != nil {
			return nil, fmt.Errorf("Error creating a Product, error: %v", err.Error())
		}

		return nil, nil
	}
}
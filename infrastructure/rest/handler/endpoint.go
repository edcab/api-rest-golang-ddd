package handler

import (
	"api-rest-golang-ddd/infrastructure/rest/model/create"
	"context"
	"encoding/json"
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

		return nil, nil
	}
}

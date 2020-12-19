package handler

import (
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	URLAdd = "/v1"
)

type Router interface {
	CreateHandler()
}

type httpHandler struct {
}

func NewHandler() *httpHandler {
	return &httpHandler{}
}

type endpoints struct {
	Add func() endpoint.Endpoint
}

func makeServerEndpoints() endpoints {
	return endpoints{
		Add: MakeCreateProductEndpoint,
	}
}

func (http httpHandler) CreateHandler() http.Handler {

	r := mux.NewRouter()
	e := makeServerEndpoints()

	r.Methods("POST").Path(URLAdd).Handler(kithttp.NewServer(
		e.Add(),
		DecodeRequest,
		EncodeResponse,
	))

	return r
}

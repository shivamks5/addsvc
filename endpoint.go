package addsvc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type AddRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type AddResponse struct {
	Result int `json:"result"`
}

func makeAddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		result := svc.Add(ctx, req.A, req.B)
		return AddResponse{Result: result}, nil
	}
}

type Endpoints struct {
	AddEndpoint endpoint.Endpoint
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		AddEndpoint: makeAddEndpoint(svc),
	}
}

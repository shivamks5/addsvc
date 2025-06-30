package addsvc

import "context"

type Service interface {
	Add(ctx context.Context, a, b int) int
}

type addService struct{}

func NewService() Service {
	return &addService{}
}

func (s *addService) Add(ctx context.Context, a, b int) int {
	return a + b
}

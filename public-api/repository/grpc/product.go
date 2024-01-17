package grpc

import (
	"be-service-public-api/domain"

	grpcpool "github.com/processout/grpc-go-pool"
)

type grpcRepositoryAuth struct {
	Pool *grpcpool.Pool
}

func NewGRPCProductRepository(Pool *grpcpool.Pool) domain.PublicAPIGRPCRepo {
	return &grpcRepositoryAuth{Pool}
}

package usecase

import (
	"be-service-public-api/domain"
)

type publicAPIUseCase struct {
	publicAPIUsecase   domain.PublicAPIUseCase
	publicAPIMySQLRepo domain.PublicAPIMySQLRepo
	publicAPIGRPCRepo  domain.PublicAPIGRPCRepo
}

func NewPublicAPIUsecase(PublicAPIMySQLRepo domain.PublicAPIMySQLRepo, PublicAPIGRPCRepo domain.PublicAPIGRPCRepo) domain.PublicAPIUseCase {
	return &publicAPIUseCase{
		publicAPIMySQLRepo: PublicAPIMySQLRepo,
		publicAPIGRPCRepo:  PublicAPIGRPCRepo,
	}
}

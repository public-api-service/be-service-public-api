package usecase

import (
	"be-service-public-api/domain"
	"context"

	serveroauth2 "gopkg.in/oauth2.v3/server"
)

type publicAPIUseCase struct {
	publicAPIUsecase   domain.PublicAPIUseCase
	publicAPIMySQLRepo domain.PublicAPIMySQLRepo
	productGRPCRepo    domain.ProductGRPCRepo
	customerGRPCRepo   domain.CustomerGRPCRepo
	oautHttp           *serveroauth2.Server
}

func NewPublicAPIUsecase(PublicAPIMySQLRepo domain.PublicAPIMySQLRepo, ProductGRPCRepo domain.ProductGRPCRepo, CustomerGRPCRepo domain.CustomerGRPCRepo, oautHttp *serveroauth2.Server) domain.PublicAPIUseCase {
	return &publicAPIUseCase{
		publicAPIMySQLRepo: PublicAPIMySQLRepo,
		productGRPCRepo:    ProductGRPCRepo,
		customerGRPCRepo:   CustomerGRPCRepo,
		oautHttp:           oautHttp,
	}
}

func (pu *publicAPIUseCase) GetAllProduct(ctx context.Context, request domain.RequestAdditionalData) (response domain.GetAllProductResponse, err error) {
	res, err := pu.productGRPCRepo.GetAllProduct(ctx, request)
	if err != nil {
		return response, err
	}

	response = res
	return
}

func (pu *publicAPIUseCase) PostCheckout(ctx context.Context, request domain.RequestDataCheckout) (err error) {
	err = pu.customerGRPCRepo.PostCheckout(ctx, request)
	if err != nil {
		return err
	}
	return
}

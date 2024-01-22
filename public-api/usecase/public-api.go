package usecase

import (
	"be-service-public-api/domain"
	"context"

	"github.com/labstack/gommon/log"
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

func (pu *publicAPIUseCase) GetProduct(ctx context.Context, request int) (response domain.ProductResponseDTO, err error) {
	response, err = pu.productGRPCRepo.GetProductByID(ctx, int64(request))
	if err != nil {
		log.Error("Error getting product data")
		return domain.ProductResponseDTO{}, err
	}

	return response, nil
}

func (pu *publicAPIUseCase) CheckStok(ctx context.Context, id int32) (err error) {
	err = pu.customerGRPCRepo.CheckStok(ctx, id)
	if err != nil {
		return err
	}
	return
}

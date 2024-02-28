package usecase

import (
	"be-service-public-api/domain"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
	serveroauth2 "gopkg.in/oauth2.v3/server"
)

type publicAPIUseCase struct {
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

func (pu *publicAPIUseCase) AccountRequest(ctx context.Context, request domain.TransactionRequest) (response domain.AdditionalFields, err error) {
	lastID, err := pu.publicAPIMySQLRepo.LastTransaction(ctx)
	if err != nil {
		log.Error(err)
		return
	}

	lastIDAI := lastID + 1

	var lastIDStr string

	if lastID >= 1000 {
		lastIDStr = "0" + strconv.Itoa(int(lastIDAI))
	}

	if lastID > 100 {
		lastIDStr = "0" + strconv.Itoa(int(lastIDAI))
	}

	if lastID >= 10 {
		lastIDStr = "00" + strconv.Itoa(int(lastIDAI))
	}

	if lastID < 10 {
		lastIDStr = "000" + strconv.Itoa(int(lastIDAI))
	}

	response.RedemptionPin = lastIDStr

	productID, err := strconv.ParseInt(request.ProductID, 10, 64)
	if err != nil {
		// Handle kesalahan jika konversi gagal
		log.Error("Error converting ProductID:", err)
		return response, err
	}
	resProduct, err := pu.productGRPCRepo.GetProductByID(ctx, productID)
	if err != nil {
		return response, err
	}

	res, err := pu.productGRPCRepo.GetListKeyProductByProductIDAndLimit(ctx, domain.RequestProductIDAndLimit{
		ProductID: request.ProductID,
		Limit:     "1",
	})

	if err != nil {
		log.Error(err)
		return response, err
	}

	var paramIDJoinStr, paramKeyNumberStr string
	for _, v := range res {
		paramIDJoinStr += strconv.Itoa(int(v.ID)) + ","
		paramKeyNumberStr += v.NumberKeys + ","
	}

	if len(paramIDJoinStr) > 0 {
		paramIDJoinStr = paramIDJoinStr[:len(paramIDJoinStr)-1]
	}

	if len(paramKeyNumberStr) > 0 {
		paramKeyNumberStr = paramKeyNumberStr[:len(paramKeyNumberStr)-1]
	}

	_, err = pu.productGRPCRepo.UpdateListKeyStatusProduct(ctx, domain.RequestUpdateKey{
		ProductID: paramIDJoinStr,
	})

	response.ActivationAccountNumber = paramKeyNumberStr
	response.BalanceAmount = strconv.Itoa(int(resProduct.FinalPrice))
	response.RedemptionAccountNumber = resProduct.SKU

	parts := strings.Fields(resProduct.Duration)
	number, _ := strconv.Atoi(parts[0])
	unit := parts[1]
	currentTime := time.Now()
	var expired time.Time

	if unit == "bulan" {
		expired = currentTime.AddDate(0, number, 0)
	} else {
		expired = currentTime.AddDate(number, 0, 0)
	}

	response.ExpiryDate = expired.Format("2006-01-02 15:04:05")

	if err != nil {
		return response, err
	}

	request.ActivationAccountNumber = paramKeyNumberStr
	request.BalanceAmount = int(resProduct.FinalPrice)
	request.RedemptionAccountNumber = resProduct.SKU
	request.ExpiryDate = expired.Format("2006-01-02 15:04:05")

	err = pu.publicAPIMySQLRepo.InsertOriginalTransaction(ctx, request)
	if err != nil {
		return response, err
	}
	return
}

func (pu *publicAPIUseCase) AccountReverse(ctx context.Context, request domain.TransactionRequest) (response domain.AdditionalFields, err error) {
	// Validation MerchantTerminalID
	err = pu.publicAPIMySQLRepo.IsExistReversalAccount(ctx, request.TransactionUniqueId)
	if err != nil {
		return response, err
	}
	err = pu.publicAPIMySQLRepo.GetDataMerchantExist(ctx, request.MerchantTerminalId)
	if err != nil {
		return response, err
	}

	// Validation MerchantTerminalIdentifier
	if request.MerchantTerminalId != request.MerchantTerminalId+"    " {
		err = errors.New("Invalid merchant identifier")
		return response, err
	}

	regex := regexp.MustCompile(`^[0-9]+$`)

	// Lakukan validasi
	if !regex.MatchString(request.LocalTransactionTime) {
		return response, errors.New("Invalid local transaction time")
	}

	if !regex.MatchString(request.LocalTransactionDate) {
		return response, errors.New("Invalid local transaction date")
	}

	productID, err := strconv.ParseInt(request.ProductID, 10, 64)
	if err != nil {
		// Handle kesalahan jika konversi gagal
		fmt.Println("Error converting ProductID:", err)
		return
	}
	resProduct, err := pu.productGRPCRepo.GetProductByID(ctx, productID)
	if err != nil {
		return response, err
	}

	res, err := pu.productGRPCRepo.GetListKeyProductByProductIDAndLimit(ctx, domain.RequestProductIDAndLimit{
		ProductID: request.ProductID,
		Limit:     "1",
	})

	if err != nil {
		log.Error(err)
		return response, err
	}

	var paramIDJoinStr, paramKeyNumberStr string
	for _, v := range res {
		paramIDJoinStr += strconv.Itoa(int(v.ID)) + ","
		paramKeyNumberStr += v.NumberKeys + ","
	}

	if len(paramIDJoinStr) > 0 {
		paramIDJoinStr = paramIDJoinStr[:len(paramIDJoinStr)-1]
	}

	if len(paramKeyNumberStr) > 0 {
		paramKeyNumberStr = paramKeyNumberStr[:len(paramKeyNumberStr)-1]
	}

	_, err = pu.productGRPCRepo.UpdateListKeyStatusProduct(ctx, domain.RequestUpdateKey{
		ProductID: paramIDJoinStr,
	})

	response.ActivationAccountNumber = paramKeyNumberStr
	response.BalanceAmount = strconv.Itoa(int(resProduct.FinalPrice))
	response.RedemptionAccountNumber = resProduct.SKU

	parts := strings.Fields(resProduct.Duration)
	number, _ := strconv.Atoi(parts[0])
	unit := parts[1]
	currentTime := time.Now()
	var expired time.Time

	if unit == "bulan" {
		expired = currentTime.AddDate(0, number, 0)
	} else {
		expired = currentTime.AddDate(number, 0, 0)
	}

	response.ExpiryDate = expired.Format("2006-01-02 15:04:05")

	if err != nil {
		return response, err
	}

	request.ActivationAccountNumber = paramKeyNumberStr
	request.BalanceAmount = int(resProduct.FinalPrice)
	request.RedemptionAccountNumber = resProduct.SKU
	request.ExpiryDate = expired.Format("2006-01-02 15:04:05")

	err = pu.publicAPIMySQLRepo.InsertOriginalTransaction(ctx, request)
	if err != nil {
		return response, err
	}
	return
}

func (pu *publicAPIUseCase) GetDataMerchantExist(ctx context.Context, merchantID string) (err error) {
	err = pu.publicAPIMySQLRepo.GetDataMerchantExist(ctx, merchantID)
	if err != nil {
		return err
	}
	return
}

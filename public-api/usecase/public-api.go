package usecase

import (
	"be-service-public-api/domain"
	"be-service-public-api/helper"
	"context"
	"errors"
	"fmt"
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

func (pu *publicAPIUseCase) AccountRequest(ctx context.Context, request domain.TransactionDTO) (response domain.AdditionalFields, err error) {
	// checkUPCAggrement := helper.CheckDataAvailabilityUPC(request.ProductID)
	// if !checkUPCAggrement {
	// 	log.Info("UPC Request is", checkUPCAggrement)
	// }

	_, err = helper.IsValidCurrencyCode(request.TransactionCurrencyCode)
	if err != nil {
		return
	}

	productID, err := strconv.ParseInt(request.ProductID, 10, 64)
	if err != nil {
		// Handle kesalahan jika konversi gagal
		log.Error("Error converting ProductID:", err)
		return response, err
	}

	resProduct, err := pu.productGRPCRepo.GetProductByID(ctx, productID)
	if err != nil {
		// if err.Error() == "Out of stok" {
		// 	if !checkUPCAggrement {
		// 		err = errors.New("Out of stok with UPC not aggrement")
		// 		return response, err
		// 	}
		// }
		return response, err
	}

	amountInt, err := strconv.Atoi(request.TransactionAmount)
	if err != nil {
		return
	}
	if resProduct.FinalPrice != float64(amountInt) {
		log.Error("Invalid amount")
		err = errors.New("Invalid amount")
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

	responseKey, err := pu.productGRPCRepo.UpdateListKeyStatusProduct(ctx, domain.RequestUpdateKey{
		ProductID: paramIDJoinStr,
		Status:    "Purchased",
	})

	if err != nil {
		log.Error("Error while update list key status ", err)
	}

	log.Info("Update list key status product :", responseKey)

	response.ActivationAccountNumber = paramKeyNumberStr
	response.BalanceAmount = strconv.Itoa(int(resProduct.FinalPrice))

	ulidString, err := helper.GenerateRedemtionAccountNumber()
	if err != nil {
		return response, err
	}
	redemptionAccountNumber := ulidString

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

	response.ExpiryDate = expired.Format("060102")

	if err != nil {
		return response, err
	}

	request.ActivationAccountNumber = redemptionAccountNumber
	request.BalanceAmount = strconv.Itoa(int(resProduct.FinalPrice))
	request.RedemptionAccountNumber = paramKeyNumberStr
	request.ExpiryDate = expired.Format("060102")
	response.ActivationAccountNumber = redemptionAccountNumber
	response.RedemptionAccountNumber = paramKeyNumberStr

	err = pu.publicAPIMySQLRepo.InsertOriginalTransaction(ctx, request)
	if err != nil {
		return response, err
	}

	err = pu.customerGRPCRepo.PostCheckoutPartner(ctx, domain.RequestDataCheckout{
		Email:            request.AcquiringInstitutionIdentifier,
		Name:             request.MerchantLocation,
		ProductSalesID:   productID,
		QTY:              1,
		TotalPricing:     int64(resProduct.FinalPrice),
		PaymentReference: request.RetrievalReferenceNumber,
		PaymentDomain:    "Blackhawk",
		ListKey:          paramKeyNumberStr,
		Invoice:          request.ProcessingCode,
		TypeDuration:     resProduct.Duration,
		Pricing:          resProduct.Price,
		Discount:         float64(*resProduct.Discount),
		Tax:              resProduct.Tax,
		Status:           "Blackhawk DAR Request",
		References:       "External",
	})

	if err != nil {
		return response, err
	}
	return
}

func (pu *publicAPIUseCase) AccountReverse(ctx context.Context, request domain.TransactionDTO) (response domain.AdditionalFields, err error) {
	resDAR, err := pu.publicAPIMySQLRepo.GetDataDigitalAccountRequest(ctx, request.RetrievalReferenceNumber)
	if err != nil {
		if err.Error() != "Not found" {
			return response, err
		}
		resDARParam, _ := pu.publicAPIMySQLRepo.GetDataDigitalAccountRequestByParam(ctx, domain.DigitalAccountReverseParam{
			ProcessingCode:                 request.ProcessingCode,
			TransactionAmount:              request.TransactionAmount,
			LocalTransactionTime:           request.LocalTransactionTime,
			LocalTransactionDate:           request.LocalTransactionDate,
			AcquiringInstitutionIdentifier: request.AcquiringInstitutionIdentifier,
			MerchantTerminalID:             request.MerchantTerminalId,
			MerchantIdentifier:             request.MerchantIdentifier,
		})
		resDAR = resDARParam

	}

	response.BalanceAmount = resDAR.TransactionAmount
	err = pu.publicAPIMySQLRepo.IsExistReversalAccount(ctx, request.RetrievalReferenceNumber)
	if err != nil {
		log.Error(err)
		err = errors.New("Duplicate Digital Account Reversal")
		return response, err
	}

	validateTransactionLocalTime := helper.ValidateTransactionTime(request.LocalTransactionTime)
	if !validateTransactionLocalTime {
		err = errors.New("Invalid transaction local time")
		return response, err
	}

	validateTransactionLocalDate := helper.ValidateLocalTransactionDate(request.LocalTransactionDate)
	if !validateTransactionLocalDate {
		err = errors.New("Invalid transaction local date")
		return response, err
	}

	if resDAR.LocalTransactionDate != request.LocalTransactionDate {
		err = errors.New("Invalid transaction local date DAR")
		return response, err
	}

	if resDAR.LocalTransactionTime != request.LocalTransactionTime {
		err = errors.New("Invalid transaction local time DAR")
		return response, err
	}

	if resDAR.TransactionAmount != request.TransactionAmount {
		err = errors.New("Invalid balance amount")
		return response, err
	}

	if resDAR.AcquiringInstitutionIdentifier != request.AcquiringInstitutionIdentifier {
		err = errors.New("Invalid Acquiring Institution Identifier")
		return response, err
	}

	if resDAR.RetrievalReferenceNumber != request.RetrievalReferenceNumber {
		err = errors.New("Invalid Retrieval Reference Number")
		return response, err
	}

	if resDAR.MerchantIdentifier != request.MerchantIdentifier {
		err = errors.New("Invalid Merchant Identifier")
		return response, err
	}

	if resDAR.MerchantTerminalId != request.MerchantTerminalId {
		err = errors.New("Invalid Merchant Terminal ID")
		return response, err
	}

	productID, err := strconv.ParseInt(request.ProductID, 10, 64)
	if err != nil {
		// Handle kesalahan jika konversi gagal
		fmt.Println("Error converting ProductID:", err)
		return
	}
	_, err = pu.productGRPCRepo.GetProductByID(ctx, productID)
	if err != nil {
		return response, err
	}

	request.BalanceAmount = resDAR.BalanceAmount
	err = pu.publicAPIMySQLRepo.InsertOriginalTransaction(ctx, request)
	if err != nil {
		return response, err
	}

	resCheckout, err := pu.customerGRPCRepo.GetCheckoutBySerialNumber(ctx, resDAR.RedemptionAccountNumber)
	if err != nil {
		return response, err
	}

	err = pu.customerGRPCRepo.PostCheckoutPartner(ctx, domain.RequestDataCheckout{
		Email:            request.AcquiringInstitutionIdentifier,
		Name:             request.MerchantLocation,
		ProductSalesID:   productID,
		QTY:              1,
		TotalPricing:     int64(resCheckout.TotalPricing),
		PaymentReference: request.RetrievalReferenceNumber,
		PaymentDomain:    "Blackhawk",
		ListKey:          resDAR.RedemptionAccountNumber,
		Invoice:          request.ProcessingCode,
		TypeDuration:     resCheckout.TypeDuration,
		Pricing:          resCheckout.Pricing,
		Discount:         resCheckout.Discount,
		Tax:              resCheckout.Tax,
		Status:           "Blackhawk DAR Reversal",
		References:       "External",
	})

	if err != nil {
		return response, err
	}

	_, err = pu.productGRPCRepo.UpdatedStatusDynamicByKeyNumber(ctx, domain.RequestUpdateKey{
		ProductID: resDAR.ActivationAccountNumber,
		Status:    "Reversal",
	})

	return
}

func (pu *publicAPIUseCase) GetDataMerchantExist(ctx context.Context, merchantID string) (err error) {
	err = pu.publicAPIMySQLRepo.GetDataMerchantExist(ctx, merchantID)
	if err != nil {
		return err
	}
	return
}

func (pu *publicAPIUseCase) InsertLog(ctx context.Context, request domain.LogRequest) (err error) {
	err = pu.publicAPIMySQLRepo.InsertLog(ctx, request)
	if err != nil {
		return err
	}
	return
}

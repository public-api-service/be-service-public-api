package handler

import (
	"be-service-public-api/domain"
	"be-service-public-api/helper"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

type PublicHandler struct {
	PublicAPIUseCase   domain.PublicAPIUseCase
	PublicAPIMySQLRepo domain.PublicAPIMySQLRepo
}

func (ph *PublicHandler) GetAllProduct(c *fiber.Ctx) error {
	var input domain.RequestAdditionalData
	search := c.Query("search")
	if search != "" {
		input.NameSearch = &search
	}

	limit := c.Query("limit")
	if limit == "" {
		limitInt := viper.GetInt("database.default_limit_query")
		log.Info("Parameter limit not set, running with default limit query from config", limitInt)

		input.Limit = &limitInt
	} else {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
		}
		input.Limit = &limitInt
	}

	page := c.Query("page")
	if page == "" {
		pageInt := viper.GetInt("database.default_page")
		log.Info("Parameter limit not set, running with default limit query from config", pageInt)

		input.Page = &pageInt
	} else {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
		}
		input.Page = &pageInt
	}

	order := c.Query("order")
	if order == "" {
		input.Order = nil
	} else {
		input.Order = &order
	}
	res, err := ph.PublicAPIUseCase.GetAllProduct(c.Context(), input)
	if err != nil {
		if err.Error() == "Not found" {
			return helper.HttpSimpleResponse(c, fasthttp.StatusNotFound)
		}
		return err
	}

	return c.Status(fasthttp.StatusOK).JSON(res)
}

func (ph *PublicHandler) PostCheckout(c *fiber.Ctx) (err error) {
	var input domain.RequestDataCheckout
	err = c.BodyParser(&input)
	if err != nil {
		log.Errorf(err.Error())
		return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	err = ph.PublicAPIUseCase.PostCheckout(c.Context(), input)
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = Stok not found" {
			return c.Status(fasthttp.StatusBadRequest).SendString("Out of stok")
		}

		if err.Error() == "rpc error: code = Unknown desc = Your transaction not found" {
			return c.Status(fasthttp.StatusBadRequest).SendString("Your transaction not found")
		}
		return err
	}

	// log.Info(input)
	return c.SendStatus(200)
}

func (ph *PublicHandler) GetProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		log.Error(err)
		return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	res, err := ph.PublicAPIUseCase.GetProduct(c.Context(), int(id))

	if err != nil {
		if err.Error() == "Data not found" {
			return c.Status(fasthttp.StatusOK).JSON(err.Error())
		}
		return helper.HttpSimpleResponse(c, fasthttp.StatusInternalServerError)

	}

	return c.Status(fasthttp.StatusOK).JSON(res)
}

func (ph *PublicHandler) CheckStok(c *fiber.Ctx) (err error) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		log.Error(err)
		return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	err = ph.PublicAPIUseCase.CheckStok(c.Context(), int32(int(id)))
	if err != nil {
		return c.Status(fasthttp.StatusNotFound).SendString("Out of stok")
	}
	return c.SendStatus(fasthttp.StatusOK)
}

// func (ph *PublicHandler) AccountRequest(c *fiber.Ctx) (err error) {
// 	var request domain.JsonRequest
// 	if err := c.BodyParser(&request); err != nil {
// 		log.Println(err)
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
// 	}

// 	res, err := ph.PublicAPIUseCase.AccountRequest(c.Context(), request)
// 	if err != nil {
// 		log.Error("Error get key : ", err)
// 		if err.Error() == "Data not found" {
// 			return helper.HttpSimpleResponse(c, fasthttp.StatusNotFound)
// 		}
// 		return err
// 	}

// 	// log.Info(res)

// 	additionalFields := domain.AdditionalFields{
// 		ActivationAccountNumber:       res.ActivationAccountNumber,
// 		BalanceAmount:                 res.BalanceAmount,
// 		CorrelatedTransactionUniqueId: request.Transaction.AdditionalTxnFields.CorrelatedTransactionUniqueId,
// 		ExpiryDate:                    res.ExpiryDate,
// 		ProductId:                     request.Transaction.AdditionalTxnFields.ProductId,
// 		RedemptionAccountNumber:       res.RedemptionAccountNumber,
// 		RedemptionPin:                 "1234",
// 		TransactionUniqueId:           request.Transaction.AdditionalTxnFields.TransactionUniqueId,
// 	}

// 	response := request
// 	response.Transaction.AdditionalTxnFields = additionalFields

// 	fullResponse := map[string]interface{}{
// 		"response": response,
// 	}
// 	return c.Status(fiber.StatusOK).JSON(fullResponse)
// }

func (ph *PublicHandler) AccountRequest(c *fiber.Ctx) (err error) {
	formattedDate := time.Now().Format("060102")
	formatTimeStamp := time.Now().Format("150405000")
	InvoiceFormat := "NRT-DN-BH"

	termAndCondition := "Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999).  Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999).  Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999).  Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions will be displayed here."

	lastInvoice := fmt.Sprintf("%s-%s-%s", InvoiceFormat, formattedDate, formatTimeStamp)

	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	jsonString, err := json.Marshal(request)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	var req domain.RequestMarshal
	err = json.Unmarshal(jsonString, &req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	productIDRegex := regexp.MustCompile(`"productId":\s*"([^"]+)"`)
	matches := productIDRegex.FindStringSubmatch(string(jsonString))
	if len(matches) != 2 {
		log.Println("Failed to extract productID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to extract productID"})
	}
	productID := matches[1]
	if transaction, ok := request["transaction"].(map[string]interface{}); ok {
		if _, exists := transaction["merchantLocation"]; exists {
			delete(transaction, "merchantLocation")
		}

		transaction["responseCode"] = "00"
	}
	res, err := ph.PublicAPIUseCase.AccountRequest(c.Context(), domain.TransactionRequest{
		ProductID:                      productID,
		Signature:                      req.Header.Signature,
		ProductCategoryCode:            req.Header.Details.ProductCategoryCode,
		SpecVersion:                    req.Header.Details.SpecVersion,
		PrimaryAccountNumber:           req.Transaction.PrimaryAccountNumber,
		ProcessingCode:                 req.Transaction.ProcessingCode,
		TransactionAmount:              req.Transaction.TransactionAmount,
		TransmissionDateTime:           req.Transaction.TransmissionDateTime,
		SystemTraceAuditNumber:         req.Transaction.SystemTraceAuditNumber,
		LocalTransactionTime:           req.Transaction.LocalTransactionTime,
		LocalTransactionDate:           req.Transaction.LocalTransactionDate,
		MerchantCategoryCode:           req.Transaction.MerchantCategoryCode,
		PointOfServiceEntryMode:        req.Transaction.PointOfServiceEntryMode,
		AcquiringInstitutionIdentifier: req.Transaction.AcquiringInstitutionIdentifier,
		RetrievalReferenceNumber:       req.Transaction.RetrievalReferenceNumber,
		MerchantTerminalId:             req.Transaction.MerchantTerminalID,
		MerchantIdentifier:             req.Transaction.MerchantIdentifier,
		MerchantLocation:               req.Transaction.MerchantLocation,
		TransactionCurrencyCode:        req.Transaction.TransactionCurrencyCode,
		TransactionUniqueId:            req.Transaction.AdditionalTxnFields.TransactionUniqueID,
		CorrelatedTransactionUniqueId:  req.Transaction.AdditionalTxnFields.CorrelatedTransactionUniqueID,
		Status:                         "Original",
	})
	if err != nil {
		responseCode := "13"
		log.Error("Error : ", err)
		log.Info(err.Error())
		if err.Error() == "rpc error: code = Unknown desc = Data not found" {
			responseCode = "00"
			request["transaction"].(map[string]interface{})["additionalTxnFields"].(map[string]interface{})["redemptionAccountNumber"] = lastInvoice

		}

		if err.Error() == "Merchant not exist" {
			responseCode = "17"
		}
		if header, ok := request["header"].(map[string]interface{}); ok {
			if details, ok := header["details"].(map[string]interface{}); ok {
				details["statusCode"] = "00"
			}
		}

		if transaction, ok := request["transaction"].(map[string]interface{}); ok {
			transaction["responseCode"] = responseCode
			transaction["termsAndConditions"] = termAndCondition
		}
		request["transaction"].(map[string]interface{})["additionalTxnFields"].(map[string]interface{})["balanceAmount"] = "C000000000000"
		return c.Status(fasthttp.StatusOK).JSON(request)
	}

	// log.Info(res)

	additionalFields := domain.AdditionalFields{
		ActivationAccountNumber: res.ActivationAccountNumber,
		BalanceAmount:           res.BalanceAmount,
		ExpiryDate:              res.ExpiryDate,
		RedemptionAccountNumber: res.RedemptionAccountNumber,
	}

	if header, ok := request["header"].(map[string]interface{}); ok {
		if details, ok := header["details"].(map[string]interface{}); ok {
			details["statusCode"] = "00"
		}
	}

	request["transaction"].(map[string]interface{})["additionalTxnFields"].(map[string]interface{})["activationAccountNumber"] = additionalFields.ActivationAccountNumber
	request["transaction"].(map[string]interface{})["additionalTxnFields"].(map[string]interface{})["balanceAmount"] = additionalFields.BalanceAmount
	request["transaction"].(map[string]interface{})["additionalTxnFields"].(map[string]interface{})["expiryDate"] = additionalFields.ExpiryDate
	request["transaction"].(map[string]interface{})["additionalTxnFields"].(map[string]interface{})["redemptionAccountNumber"] = additionalFields.RedemptionAccountNumber

	return c.Status(fiber.StatusOK).JSON(request)
}

func (ph *PublicHandler) AccountReverse(c *fiber.Ctx) (err error) {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	jsonString, err := json.Marshal(request)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	productIDRegex := regexp.MustCompile(`"productId":\s*"([^"]+)"`)
	matches := productIDRegex.FindStringSubmatch(string(jsonString))
	if len(matches) != 2 {
		log.Println("Failed to extract productID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to extract productID"})
	}
	productID := matches[1]

	transaction := request["transaction"].(map[string]interface{})
	additionalTxnFields := transaction["additionalTxnFields"].(map[string]interface{})
	res, err := ph.PublicAPIUseCase.AccountReverse(c.Context(), domain.TransactionRequest{
		ProductID:            productID,
		LocalTransactionDate: transaction["transmissionDateTime"].(string),
		LocalTransactionTime: transaction["localTransactionTime"].(string),
		MerchantTerminalId:   transaction["merchantTerminalId"].(string),
		MerchantIdentifier:   transaction["merchantIdentifier"].(string),
		TransactionUniqueId:  additionalTxnFields["transactionUniqueId"].(string),
	})
	if err != nil {
		if header, ok := request["header"].(map[string]interface{}); ok {
			if details, ok := header["details"].(map[string]interface{}); ok {
				details["statusCode"] = "00"
			}
		}
		responseCode := "12"
		termAndCondition := "Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999).  Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999).  Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999).  Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions will be displayed here."

		log.Error("Err usecase  : ", err)
		if err.Error() == "Data not found" {
			responseCode = "16"
		}

		if err.Error() == "Merchant not exist" || err.Error() == "Invalid merchant identifier" {
			responseCode = "12"
		}

		if err.Error() == "Duplicate reversal account" {
			responseCode = "34"
		}

		if transaction, ok := request["transaction"].(map[string]interface{}); ok {
			transaction["responseCode"] = responseCode
			transaction["termsAndConditions"] = termAndCondition
		}

		return c.Status(fasthttp.StatusOK).JSON(request)
	}

	// log.Info(res)

	if transaction, ok := request["transaction"].(map[string]interface{}); ok {
		if _, exists := transaction["merchantLocation"]; exists {
			delete(transaction, "merchantLocation")
		}
	}
	additionalFields := domain.AdditionalFields{
		ActivationAccountNumber: res.ActivationAccountNumber,
		BalanceAmount:           res.BalanceAmount,
		ExpiryDate:              res.ExpiryDate,
		RedemptionAccountNumber: res.RedemptionAccountNumber,
	}

	request["transaction"].(map[string]interface{})["additionalTxnFields"].(map[string]interface{})["balanceAmount"] = additionalFields.BalanceAmount

	// Add additional fields to the "header" section
	if header, ok := request["header"].(map[string]interface{}); ok {
		if details, ok := header["details"].(map[string]interface{}); ok {
			details["statusCode"] = "00"
		}
	}

	if transaction, ok := transaction["transaction"].(map[string]interface{}); ok {
		transaction["responseCode"] = "00"
	}
	return c.Status(fiber.StatusOK).JSON(request)
}

func (ph *PublicHandler) Network(c *fiber.Ctx) (err error) {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	// Memastikan bahwa request memiliki struktur yang diharapkan
	reqData, ok := request["request"].(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request structure"})
	}

	// Mengatur status code pada bagian header jika struktur request sesuai
	if header, ok := reqData["header"].(map[string]interface{}); ok {
		if details, ok := header["details"].(map[string]interface{}); ok {
			details["statusCode"] = "00"
		}
	}

	// Mengatur nilai-nilai pada bagian transaction jika struktur request sesuai
	if transaction, ok := reqData["transaction"].(map[string]interface{}); ok {
		transaction["authIdentificationResponse"] = "000000"
		transaction["responseCode"] = "00"
	}

	// Membuat respons dengan struktur yang sesuai
	response := map[string]interface{}{
		"response": reqData,
	}

	return c.Status(fasthttp.StatusOK).JSON(response)
}

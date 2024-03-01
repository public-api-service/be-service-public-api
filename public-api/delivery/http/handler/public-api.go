package handler

import (
	"be-service-public-api/domain"
	"be-service-public-api/helper"
	"encoding/json"
	"strconv"

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
// 	formattedDate := time.Now().Format("060102")
// 	formatTimeStamp := time.Now().Format("150405000")
// 	InvoiceFormat := "NRT-DN-BH"

// 	termAndCondition := "Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999).  Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999).  Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999).  Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions will be displayed here."

// 	lastInvoice := fmt.Sprintf("%s-%s-%s", InvoiceFormat, formattedDate, formatTimeStamp)

// 	var request map[string]interface{}
// 	if err := c.BodyParser(&request); err != nil {
// 		log.Println(err)
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
// 	}

// 	jsonString, err := json.Marshal(request)
// 	if err != nil {
// 		log.Println(err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
// 	}
// 	productIDRegex := regexp.MustCompile(`"productId":\s*"([^"]+)"`)
// 	matches := productIDRegex.FindStringSubmatch(string(jsonString))
// 	if len(matches) != 2 {
// 		log.Println("Failed to extract productID")
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to extract productID"})
// 	}
// 	productID := matches[1]

// 	regex := `\"request\"\s*:\s*\{`
// 	re := regexp.MustCompile(regex)

// 	var transaction, header, detail map[string]interface{}
// 	if re.MatchString(string(jsonString)) {
// 		log.Info("Request menggunakan param request")
// 		requestJSON := request["request"].(map[string]interface{})
// 		transaction = requestJSON["transaction"].(map[string]interface{})
// 		header = requestJSON["header"].(map[string]interface{})
// 		detail = header["details"].(map[string]interface{})

// 	} else {
// 		log.Info("Request tidak menggunakan param request")
// 		transaction = request["transaction"].(map[string]interface{})
// 		header = request["header"].(map[string]interface{})
// 		detail = header["details"].(map[string]interface{})
// 	}

// 	res, err := ph.PublicAPIUseCase.AccountRequest(c.Context(), domain.TransactionDTO{
// 		ProductID:                      productID,
// 		Signature:                      header["signature"].(string),
// 		ProductCategoryCode:            detail["productCategoryCode"].(string),
// 		SpecVersion:                    detail["specVersion"].(string),
// 		PrimaryAccountNumber:           transaction["primaryAccountNumber"].(string),
// 		ProcessingCode:                 transaction["processingCode"].(string),
// 		TransactionAmount:              transaction["transactionAmount"].(string),
// 		TransmissionDateTime:           transaction["transmissionDateTime"].(string),
// 		SystemTraceAuditNumber:         transaction["systemTraceAuditNumber"].(string),
// 		LocalTransactionTime:           transaction["localTransactionTime"].(string),
// 		LocalTransactionDate:           transaction["localTransactionDate"].(string),
// 		MerchantCategoryCode:           transaction["merchantCategoryCode"].(string),
// 		PointOfServiceEntryMode:        transaction["pointOfServiceEntryMode"].(string),
// 		AcquiringInstitutionIdentifier: transaction["acquiringInstitutionIdentifier"].(string),
// 		RetrievalReferenceNumber:       transaction["retrievalReferenceNumber"].(string),
// 		MerchantTerminalId:             transaction["merchantTerminalId"].(string),
// 		MerchantIdentifier:             transaction["merchantIdentifier"].(string),
// 		MerchantLocation:               transaction["merchantLocation"].(string),
// 		TransactionCurrencyCode:        transaction["transactionCurrencyCode"].(string),
// 		TransactionUniqueId:            transaction["additionalTxnFields"].(map[string]interface{})["transactionUniqueId"].(string),
// 		CorrelatedTransactionUniqueId:  transaction["additionalTxnFields"].(map[string]interface{})["correlatedTransactionUniqueId"].(string),
// 		Status:                         "Original",
// 	})
// 	if err != nil {
// 		delete(transaction, "merchantLocation")
// 		transaction["responseCode"] = "00"
// 		responseCode := "13"
// 		log.Error("Error : ", err)
// 		log.Info(err.Error())
// 		if err.Error() == "rpc error: code = Unknown desc = Data not found" {
// 			responseCode = "00"
// 			request["transaction"].(map[string]interface{})["additionalTxnFields"].(map[string]interface{})["redemptionAccountNumber"] = lastInvoice

// 		}

// 		if err.Error() == "Merchant not exist" {
// 			responseCode = "17"
// 		}

// 		detail["statusCode"] = "00"

// 		transaction["responseCode"] = responseCode
// 		transaction["termsAndConditions"] = termAndCondition
// 		transaction["additionalTxnFields"].(map[string]interface{})["balanceAmount"] = "C000000000000"

// 		msgTopic := ":bangbang: :bangbang: **PUBLIC API ERROR DIGITAL TRANSACTION** :bangbang: :bangbang:  \n\n ***request : *** \n " + string(jsonString) + "\n\n ***error log : *** \n" + err.Error()
// 		err = helper.SendMessageToDiscord("https://discord.com/api/webhooks/1210122017584447499/fAXy7V14dtHULFkvWtOmjNH65sMOsve2bDW90BtYbyFVNuudy-3lNE_qFAKmkvjlJ2wH", msgTopic)
// 		if err != nil {
// 			return nil
// 		}

// 		return c.Status(fasthttp.StatusOK).JSON(request)
// 	}

// 	// log.Info(res)

// 	additionalFields := domain.AdditionalFields{
// 		ActivationAccountNumber: res.ActivationAccountNumber,
// 		BalanceAmount:           res.BalanceAmount,
// 		ExpiryDate:              res.ExpiryDate,
// 		RedemptionAccountNumber: res.RedemptionAccountNumber,
// 	}

// 	if header, ok := request["header"].(map[string]interface{}); ok {
// 		if details, ok := header["details"].(map[string]interface{}); ok {
// 			details["statusCode"] = "00"
// 		}
// 	}

// 	request["transaction"].(map[string]interface{})["additionalTxnFields"].(map[string]interface{})["activationAccountNumber"] = additionalFields.ActivationAccountNumber
// 	request["transaction"].(map[string]interface{})["additionalTxnFields"].(map[string]interface{})["balanceAmount"] = additionalFields.BalanceAmount
// 	request["transaction"].(map[string]interface{})["additionalTxnFields"].(map[string]interface{})["expiryDate"] = additionalFields.ExpiryDate
// 	request["transaction"].(map[string]interface{})["additionalTxnFields"].(map[string]interface{})["redemptionAccountNumber"] = additionalFields.RedemptionAccountNumber
// 	delete(transaction, "merchantLocation")
// 	transaction["responseCode"] = "00"
// 	return c.Status(fiber.StatusOK).JSON(request)
// }

func (ph *PublicHandler) AccountRequest(c *fiber.Ctx) (err error) {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	// Procesed JSON Request
	requestJSON := request["request"].(map[string]interface{})
	header := requestJSON["header"].(map[string]interface{})
	details := header["details"].(map[string]interface{})
	transaction := requestJSON["transaction"].(map[string]interface{})
	additionalTxnFields := transaction["additionalTxnFields"].(map[string]interface{})

	res, err := ph.PublicAPIUseCase.AccountRequest(c.Context(), domain.TransactionDTO{
		ProductID:                      additionalTxnFields["productId"].(string),
		Signature:                      header["signature"].(string),
		ProductCategoryCode:            details["productCategoryCode"].(string),
		SpecVersion:                    details["specVersion"].(string),
		PrimaryAccountNumber:           transaction["primaryAccountNumber"].(string),
		ProcessingCode:                 transaction["processingCode"].(string),
		TransactionAmount:              transaction["transactionAmount"].(string),
		TransmissionDateTime:           transaction["transmissionDateTime"].(string),
		SystemTraceAuditNumber:         transaction["systemTraceAuditNumber"].(string),
		LocalTransactionTime:           transaction["localTransactionTime"].(string),
		LocalTransactionDate:           transaction["localTransactionDate"].(string),
		MerchantCategoryCode:           transaction["merchantCategoryCode"].(string),
		PointOfServiceEntryMode:        transaction["pointOfServiceEntryMode"].(string),
		AcquiringInstitutionIdentifier: transaction["acquiringInstitutionIdentifier"].(string),
		RetrievalReferenceNumber:       transaction["retrievalReferenceNumber"].(string),
		MerchantTerminalId:             transaction["merchantTerminalId"].(string),
		MerchantIdentifier:             transaction["merchantIdentifier"].(string),
		MerchantLocation:               transaction["merchantLocation"].(string),
		TransactionCurrencyCode:        transaction["transactionCurrencyCode"].(string),
		TransactionUniqueId:            additionalTxnFields["transactionUniqueId"].(string),
		CorrelatedTransactionUniqueId:  additionalTxnFields["correlatedTransactionUniqueId"].(string),
		Status:                         "Digital Account Request",
	})

	details["statusCode"] = "00"
	transaction["responseCode"] = "00"
	transaction["authIdentificationResponse"] = "123456"
	delete(transaction, "merchantLocation")
	responseJSON := make(map[string]interface{})
	if err != nil {
		transaction["responseCode"] = "14"
		if err.Error() == "Invalid amount" {
			transaction["responseCode"] = "13"
		}
		log.Error(err)
		transaction["authIdentificationResponse"] = "000000"
		additionalTxnFields["balanceAmount"] = "C000000000000"
		// transaction["termsAndConditions"] = "Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999).  Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999).  Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999).  Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety nine (999). Terms and Conditions will be displayed here."
		responseJSON["response"] = request["request"]
		jsonString, errr := json.Marshal(responseJSON)
		if errr != nil {
			log.Error(err)
			return c.Status(fasthttp.StatusOK).JSON(responseJSON)
		}

		log.Info("Try to send log to discord")
		msgTopic := ":bangbang: :bangbang: **PUBLIC API ERROR DIGITAL TRANSACTION** :bangbang: :bangbang:  \n\n" + "***error log : *** \n" + err.Error() + "\n\n ***request : *** \n " + string(jsonString)
		errrr := helper.SendMessageToDiscord("https://discord.com/api/webhooks/1210122017584447499/fAXy7V14dtHULFkvWtOmjNH65sMOsve2bDW90BtYbyFVNuudy-3lNE_qFAKmkvjlJ2wH", msgTopic)
		if errrr != nil {
			log.Error(err)
			return nil
		}
		return c.Status(fasthttp.StatusOK).JSON(responseJSON)
	}

	transaction["authIdentificationResponse"] = helper.GenerateRandomNumber(6)
	additionalTxnFields["activationAccountNumber"] = res.ActivationAccountNumber
	additionalTxnFields["balanceAmount"] = "C" + transaction["transactionAmount"].(string)
	additionalTxnFields["redemptionAccountNumber"] = res.RedemptionAccountNumber
	additionalTxnFields["redemptionPin"] = res.RedemptionPin
	additionalTxnFields["expiryDate"] = res.ExpiryDate

	responseJSON["response"] = request["request"]

	return c.JSON(responseJSON)
}

func (ph *PublicHandler) AccountReverse(c *fiber.Ctx) (err error) {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	// Procesed JSON Request
	requestJSON := request["request"].(map[string]interface{})
	header := requestJSON["header"].(map[string]interface{})
	details := header["details"].(map[string]interface{})
	transaction := requestJSON["transaction"].(map[string]interface{})
	additionalTxnFields := transaction["additionalTxnFields"].(map[string]interface{})

	res, err := ph.PublicAPIUseCase.AccountReverse(c.Context(), domain.TransactionDTO{
		ProductID:                      additionalTxnFields["productId"].(string),
		Signature:                      header["signature"].(string),
		ProductCategoryCode:            details["productCategoryCode"].(string),
		SpecVersion:                    details["specVersion"].(string),
		PrimaryAccountNumber:           transaction["primaryAccountNumber"].(string),
		ProcessingCode:                 transaction["processingCode"].(string),
		TransactionAmount:              transaction["transactionAmount"].(string),
		TransmissionDateTime:           transaction["transmissionDateTime"].(string),
		SystemTraceAuditNumber:         transaction["systemTraceAuditNumber"].(string),
		LocalTransactionTime:           transaction["localTransactionTime"].(string),
		LocalTransactionDate:           transaction["localTransactionDate"].(string),
		MerchantCategoryCode:           transaction["merchantCategoryCode"].(string),
		PointOfServiceEntryMode:        transaction["pointOfServiceEntryMode"].(string),
		AcquiringInstitutionIdentifier: transaction["acquiringInstitutionIdentifier"].(string),
		RetrievalReferenceNumber:       transaction["retrievalReferenceNumber"].(string),
		MerchantTerminalId:             transaction["merchantTerminalId"].(string),
		MerchantIdentifier:             transaction["merchantIdentifier"].(string),
		MerchantLocation:               transaction["merchantLocation"].(string),
		TransactionCurrencyCode:        transaction["transactionCurrencyCode"].(string),
		TransactionUniqueId:            additionalTxnFields["transactionUniqueId"].(string),
		CorrelatedTransactionUniqueId:  additionalTxnFields["correlatedTransactionUniqueId"].(string),
		Status:                         "Digital Account Reverse",
	})

	details["statusCode"] = "00"
	transaction["responseCode"] = "00"
	transaction["authIdentificationResponse"] = "123456"
	delete(transaction, "merchantLocation")
	responseJSON := make(map[string]interface{})
	if err != nil {

		additionalTxnFields["balanceAmount"] = "C000000000000"
		if err.Error() == "Duplicate Digital Account Reversal" {
			transaction["responseCode"] = "34"
		}

		if err.Error() == "Invalid transaction local time" || err.Error() == "Invalid transaction local date" {
			additionalTxnFields["balanceAmount"] = "C" + res.BalanceAmount
			transaction["responseCode"] = "12"
		}

		if err.Error() == "Invalid transaction local time DAR" || err.Error() == "Invalid transaction local date DAR" {
			additionalTxnFields["balanceAmount"] = "C" + res.BalanceAmount
			transaction["responseCode"] = "12"
		}

		log.Error(err)
		transaction["authIdentificationResponse"] = "000000"
		responseJSON["response"] = request["request"]
		jsonString, errr := json.Marshal(responseJSON)
		if errr != nil {
			log.Error(err)
			return c.Status(fasthttp.StatusOK).JSON(responseJSON)
		}

		log.Info("Try to send log to discord")
		msgTopic := ":bangbang: :bangbang: **PUBLIC API ERROR DIGITAL TRANSACTION REVERSE** :bangbang: :bangbang:  \n\n" + "***error log : *** \n" + err.Error() + "\n\n ***request : *** \n " + string(jsonString)
		errrr := helper.SendMessageToDiscord("https://discord.com/api/webhooks/1210122017584447499/fAXy7V14dtHULFkvWtOmjNH65sMOsve2bDW90BtYbyFVNuudy-3lNE_qFAKmkvjlJ2wH", msgTopic)
		if errrr != nil {
			log.Error(err)
			return nil
		}
		return c.Status(fasthttp.StatusOK).JSON(responseJSON)
	}

	transaction["authIdentificationResponse"] = helper.GenerateRandomNumber(6)
	additionalTxnFields["balanceAmount"] = "C000000000000"

	responseJSON["response"] = request["request"]

	return c.JSON(responseJSON)
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

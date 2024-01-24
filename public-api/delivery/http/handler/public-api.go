package handler

import (
	"be-service-public-api/domain"
	"be-service-public-api/helper"
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
func (ph *PublicHandler) BlackHawk(c *fiber.Ctx) (err error) {
	response := domain.ResponseBlackHawk{
		Header: domain.ResponseHeaderDetailBlackHawk{
			Detail: domain.ResponseHeaderContentBlackHawk{
				ProductCategoryCode: "01",
				SpecVersion:         "46",
				StatusCode:          "00",
			},
			Signature: "BHNUMS",
		},

		Transaction: domain.ResponseTransactionBlackHawk{
			AcquiringInstitutionIdentifier: "60300000063",
			AdditionalTxnFields: domain.ResponseAdditionalTxnFieldsTransactionBlackHawk{
				ActivationAccountNumber:       "6039537201000000000",
				BalanceAmount:                 "C000000002500",
				CorrelatedTransactionUniqueId: "9WKNNBT0QBGTWWNW0DJBSPRYZ4",
				ExpiryDate:                    "491201",
				PaymentDetails: domain.ResponsePaymentBlackHawk{
					PaymentDetail: domain.ResponsePaymentDetailBlackHawk{
						PaymentMode: "051",
						TenderType:  "Credit Card",
					},
				},
				ProductId:               "07675004390",
				RedemptionAccountNumber: "XXBNC5HR7ZPN43GQ",
				RedemptionPin:           "1234      ",
				TransactionUniqueId:     "9WKNNBT0QBGTWWNW0DJBSPRYZ4",
			},
			AuthIdentificationResponse: "123456",
			LocalTransactionDate:       "230414",
			LocalTransactionTime:       "082515",
			MerchantCategoryCode:       "5411",
			MerchantIdentifier:         "60300000063    ",
			MerchantTerminalId:         "06220     900   ",
			PointOfServiceEntryMode:    "011",
			PrimaryAccountNumber:       "9877890000000000",
			ProcessingCode:             "745400",
			ReceiptsFields: domain.ResponseReceiptsFieldsBlackHawk{
				Lines: []string{
					"StoreId : 06220",
					"Address : BLACKHAWK SIM-jlee126",
					"City : PLEASANTON CA",
					"State :  US",
					"LocalTxnDate : 04/14/23",
					"LocalTxnTime : 082515",
					"Denomination : 25.00",
					"PINNumber : XXBNC5HR7ZPN43GQ",
					"PhoneNumber : 1-888-BHN-HELP",
					"SequenceNumber : 000000661586",
					"AccountBalance : 25.00",
					"ShortGUID : BSPRYZ4",
					"RedemptionPIN : 1234",
					"ActivationNum : 6039537201000000000",
					"DigitalExpDate : 491201",
					"AddtnlData :",
				},
			},
			ResponseCode:             "00",
			RetrievalReferenceNumber: "000000661586",
			SystemTraceAuditNumber:   "499180",
			TermsAndConditions:       "Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety-nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety-nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety-nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety-nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety-nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety-nine (999). Terms and Conditions of the card will be displayed in this area. The maximum characters allowed are nine hundred and ninety-nine (999). Terms and Conditions will be displayed here.",
			TransactionAmount:        "000000002500",
			TransactionCurrencyCode:  "840",
			TransmissionDateTime:     "230414082515",
		},
	}

	fullResponse := map[string]interface{}{
		"response": response,
	}
	return c.Status(fiber.StatusOK).JSON(fullResponse)
}

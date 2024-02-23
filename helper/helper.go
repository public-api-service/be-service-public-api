package helper

import (
	"be-service-public-api/domain"
	"context"
	"encoding/base64"
	"errors"
	"regexp"
	"strconv"
	"strings"

	nethttp "net/http"

	"github.com/labstack/gommon/log"
	"github.com/rmg/iso4217"
)

func ToAlphaString(col int) string {
	var result string
	for col > 0 {
		col--
		result = string('A'+col%26) + result
		col /= 26
	}
	return result
}

func RecachingB2BDataClient(usecase domain.AuthorizationMySQLRepo) (response []domain.ResponseB2BDTO, err error) {
	response, err = usecase.GetAllClientData(context.Background())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info(response)
	return response, nil
}

func GenerateOAuthCredential(ctx context.Context, r *nethttp.Request) (clientCredential domain.OAuth2ClientCredential, err error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {

		log.Error("Authorization header missing")
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Basic" {
		log.Error("Invalid Authorization header format")
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		log.Error(err)
		return
	}

	credentials := strings.SplitN(string(decoded), ":", 2)
	if len(credentials) != 2 {
		log.Error("Invalid credentials format")
		return
	}

	log.Info("ClientID :", credentials[0])
	log.Info("ClientSecret : ", credentials[1])

	clientCredential = domain.OAuth2ClientCredential{
		ClientID:     credentials[0],
		ClientSecret: credentials[1],
	}

	return

}

func IsValidCurrencyCode(code string) (currencyCode string, err error) {
	codeInt, err := strconv.Atoi(code)
	if err != nil {
		return currencyCode, err
	}

	currencyCode, minor := iso4217.ByCode(codeInt)

	if currencyCode == "" && minor == 0 {
		return currencyCode, errors.New("Currency code invalid")
	}

	return
}

func IsValidAmount(amount, currencyCode string) error {
	msg := "Invalid amount"

	regex := regexp.MustCompile(`^[0-9]+$`)
	if !regex.MatchString(amount) {
		return errors.New(msg)
	}

	return nil
}

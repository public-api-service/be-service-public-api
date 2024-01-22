package helper

import (
	"be-service-public-api/domain"
	"context"
	"encoding/base64"
	"strings"

	nethttp "net/http"

	"github.com/labstack/gommon/log"
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

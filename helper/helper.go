package helper

import (
	"be-service-public-api/domain"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"net/http"
	nethttp "net/http"

	"github.com/labstack/gommon/log"
	"github.com/rmg/iso4217"
)

func ToAlphaString(col int) string {
	var result string
	for col > 0 {
		col--
		result = fmt.Sprintf("%c", 'A'+col%26) + result
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

func SendMessageToDiscord(webhookURL string, message string) (err error) {
	// Membuat payload pesan dalam format JSON
	message = message[:2000]
	payload := map[string]string{"content": message}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Error("error marshalling JSON payload: ", err)
		return nil
	}

	// Mengirimkan POST request ke webhook URL
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Error("error sending POST request:", err)
		return nil
	}
	defer resp.Body.Close()

	// Mengecek status code dari response
	if resp.StatusCode != http.StatusOK {
		log.Error("unexpected status code: ", resp.StatusCode)
		return nil
	}

	return nil
}

func GenerateRandomNumber(length int) (response string) {
	characters := "0123456789"

	// Buat string acak dengan panjang yang diinginkan
	result := make([]byte, length)
	for i := range result {
		result[i] = characters[rand.Intn(len(characters))]
	}
	return string(result)
}

func GenerateRandomString(length int) (response string) {
	characters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Buat string acak dengan panjang yang diinginkan
	result := make([]byte, length)
	for i := range result {
		result[i] = characters[rand.Intn(len(characters))]
	}
	return string(result)
}

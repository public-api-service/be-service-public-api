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
	"time"

	"github.com/oklog/ulid"

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

	if currencyCode != "IDR" {
		return currencyCode, errors.New("Currency code invalid")
	}

	return
}

func IsValidAmount(amount string, currencyCode string) (response float64, err error) {
	// Mendapatkan jumlah digit dari unit minor mata uang
	code, err := strconv.Atoi(currencyCode)
	if err != nil {
		return response, errors.New("invalid currency code")
	}
	_, minorDigits := iso4217.ByCode(code)

	// Mengonversi jumlah digit unit minor mata uang menjadi faktor pembulatan
	factor := 1
	for i := 0; i < minorDigits; i++ {
		factor *= 10
	}

	// Mengonversi nilai amount menjadi integer
	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		return response, errors.New("invalid amount")
	}

	// Menghitung nilai transactionAmount berdasarkan faktor pembulatan
	roundedAmount := amountInt / factor

	// Mengonversi nilai transactionAmount menjadi string dengan zero-padding
	transactionAmount := fmt.Sprintf("%0*d", minorDigits, roundedAmount)

	fmt.Println("Transaction Amount:", transactionAmount)
	return float64(roundedAmount), nil
}

func SendMessageToDiscord(webhookURL string, message string) (err error) {
	// Membuat payload pesan dalam format JSON
	if len(message) > 2000 {
		message = message[:2000]
	}
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

func ValidateTransactionTime(transactionTime string) bool {
	// Periksa panjang string
	if len(transactionTime) != 6 {
		return false
	}

	// Periksa format waktu menggunakan ekspresi reguler
	match, _ := regexp.MatchString(`^\d{6}$`, transactionTime)
	if !match {
		return false
	}

	// Konversi bagian waktu menjadi integer untuk memeriksa rentang
	hour, _ := strconv.Atoi(transactionTime[:2])
	minute, _ := strconv.Atoi(transactionTime[2:4])
	second, _ := strconv.Atoi(transactionTime[4:6])

	// Periksa rentang waktu
	if hour < 0 || hour > 23 || minute < 0 || minute > 59 || second < 0 || second > 59 {
		return false
	}

	return true
}

func ValidateLocalTransactionDate(localTransactionDate string) bool {
	// Periksa panjang string
	if len(localTransactionDate) != 6 {
		return false
	}

	// Periksa format tanggal menggunakan ekspresi reguler
	match, _ := regexp.MatchString(`^\d{6}$`, localTransactionDate)
	if !match {
		return false
	}

	// Konversi bagian tanggal menjadi integer untuk memeriksa rentang
	year, _ := strconv.Atoi(localTransactionDate[:2])
	month, _ := strconv.Atoi(localTransactionDate[2:4])
	day, _ := strconv.Atoi(localTransactionDate[4:6])

	// Periksa rentang tanggal
	if year < 0 || year > 99 || month < 1 || month > 12 || day < 1 || day > 31 {
		return false
	}

	return true
}

func GenerateRedemtionAccountNumber() (string, error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	ulid := ulid.MustNew(ulid.Timestamp(t), entropy)
	ulidString := ulid.String()
	return ulidString, nil
}

// func CheckDataAvailabilityUPC(input string) bool {
// 	dataAggrementUPC := viper.GetStringSlice("aggrement_upc")
// 	for _, value := range dataAggrementUPC {
// 		if value == input {
// 			return true
// 		}
// 	}
// 	return false
// }

package handler

import (
	"be-service-public-api/domain"
	"be-service-public-api/helper"
	"context"
	"io"
	"strings"

	nethttp "net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	log "github.com/sirupsen/logrus"
)

type AuthorizationHandler struct {
	AuthorizationUseCase   domain.AuthorizationUseCase
	AuthorizationMySQLRepo domain.AuthorizationMySQLRepo
}

func (au *AuthorizationHandler) PostTokenOAuth2(w nethttp.ResponseWriter, r *nethttp.Request) {
	var ctx context.Context
	// Generate Credential OAUTH2
	CredentialValues, _ := helper.GenerateOAuthCredential(ctx, r)

	log.Info(CredentialValues)

	// Data baru yang ingin ditambahkan ke tubuh permintaan
	newPayloadData := "&client_id=" + CredentialValues.ClientID + "&client_secret=" + CredentialValues.ClientSecret

	existingBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Error reading existing request body:", err)
		return
	}
	// r.Body.Close()

	// Menggabungkan data lama dan data baru
	newRequestBody := append(existingBody, []byte(newPayloadData)...)

	// Membuat ulang tubuh permintaan dengan data yang diperbarui
	r.Body = io.NopCloser(strings.NewReader(string(newRequestBody)))

	responseToken := au.AuthorizationUseCase.TokenOAuth(ctx, w, r)
	if responseToken != nil {
		log.Error(responseToken.Error())
	}
}

func (au *AuthorizationHandler) TokenOauth() fiber.Handler {
	return func(c *fiber.Ctx) error {

		httpReq, err := adaptor.ConvertRequest(c, false)
		if err != nil {
			return err
		}

		err = au.AuthorizationUseCase.ValidateBarrerToken(c.Context(), httpReq)

		log.Print(err)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		c.Locals("isPartner", true)

		return c.Next()
	}
}

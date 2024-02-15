package http

import (
	"be-service-public-api/public-api/delivery/http/handler"
	// "be-service-public-api/delivery/http/handler"
	"be-service-public-api/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// RouterAPI is the main router for this Service Insurance REST API
func RouterAPI(app *fiber.App, PublicAPIUseCase domain.PublicAPIUseCase, AuthorizationUseCase domain.AuthorizationUseCase) {
	handlerPublicAPI := &handler.PublicHandler{PublicAPIUseCase: PublicAPIUseCase}
	handlerAuthorization := &handler.AuthorizationHandler{AuthorizationUseCase: AuthorizationUseCase}

	basePath := viper.GetString("server.base_path")

	publicAPI := app.Group(basePath)

	publicAPI.Use(cors.New(cors.Config{
		AllowOrigins: viper.GetString("middleware.allows_origin"),
	}))

	log.Info(handlerPublicAPI)
	// Public API Route
	publicAPI.Get("/product", handlerAuthorization.TokenOauth(), handlerPublicAPI.GetAllProduct)
	publicAPI.Post("/checkout", handlerAuthorization.TokenOauth(), handlerPublicAPI.PostCheckout)
	publicAPI.Get("/product/:id", handlerAuthorization.TokenOauth(), handlerPublicAPI.GetProduct)
	publicAPI.Post("b2b/token", adaptor.HTTPHandlerFunc(handlerAuthorization.PostTokenOAuth2))
	publicAPI.Get("/check/stok/:id", handlerAuthorization.TokenOauth(), handlerPublicAPI.CheckStok)
	publicAPI.Post("/transaction", handlerPublicAPI.AccountRequest)
	publicAPI.Post("/transaction/reverse", handlerPublicAPI.AccountReverse)
	publicAPI.Post("/transaction/network", handlerPublicAPI.Network)
	// publicAPI.Get("/transaction/network", handlerPublicAPI.Network)

}

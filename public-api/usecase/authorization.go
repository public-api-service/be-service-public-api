package usecase

import (
	"be-service-public-api/domain"
	"context"
	nethttp "net/http"

	"github.com/labstack/gommon/log"
	serveroauth2 "gopkg.in/oauth2.v3/server"
)

type authorizationUseCase struct {
	authorizationUsecase   domain.AuthorizationUseCase
	authorizationMySQLRepo domain.AuthorizationMySQLRepo
	oautHttp               *serveroauth2.Server
}

func NewAuthorizationUsecase(AuthorizationMySQLRepo domain.AuthorizationMySQLRepo, oautHttp *serveroauth2.Server) domain.AuthorizationUseCase {
	return &authorizationUseCase{
		authorizationMySQLRepo: AuthorizationMySQLRepo,
		oautHttp:               oautHttp,
	}
}

func (au *authorizationUseCase) TokenOAuth(ctx context.Context, w nethttp.ResponseWriter, r *nethttp.Request) (err error) {
	err = au.oautHttp.HandleTokenRequest(w, r)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (au *authorizationUseCase) ValidateBarrerToken(ctx context.Context, r *nethttp.Request) (err error) {
	tokenInfo, err := au.oautHttp.ValidationBearerToken(r)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info(tokenInfo.GetClientID())
	return
}

func (au *authorizationUseCase) GetAllClientData(ctx context.Context) (response []domain.ResponseB2BDTO, err error) {
	response, err = au.authorizationUsecase.GetAllClientData(ctx)
	if err != nil {
		return
	}
	return
}

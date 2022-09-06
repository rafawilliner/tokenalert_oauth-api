package services

import (
	"tokenalert_oauth-api/src/domain/access_token"
	"tokenalert_oauth-api/src/repositories/db"
	"tokenalert_oauth-api/src/repositories/rest"
	"github.com/rafawilliner/tokenalert_utils-go/src/rest_errors"
)

var (
	AccessTokenService accessTokenServiceInterface = &accessTokenService{}
)

type accessTokenService struct{}

type accessTokenServiceInterface interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestErr)	
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, rest_errors.RestErr)	
}

func (a *accessTokenService) GetById(access_token_id string) (*access_token.AccessToken, rest_errors.RestErr) { 

	var at *access_token.AccessToken
	var err rest_errors.RestErr
	if at, err = db.AccessTokenRepository.GetById(access_token_id); err != nil {
		return nil, err
	}
	
	return at, nil
}

func (s *accessTokenService) Create(atr access_token.AccessTokenRequest) (*access_token.AccessToken, rest_errors.RestErr) {
	
	if err := atr.Validate(); err != nil {
		return nil, err
	}

	user, err := rest.RestUsersRepository.LoginUser(atr.Email, atr.Password)
	if err != nil {
		return nil, err
	}

	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	if err := db.AccessTokenRepository.Create(atr); err != nil {
		return nil, err
	}

	return &at, nil
}
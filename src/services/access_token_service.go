package services

import (
	"tokenalert_oauth-api/src/domain/access_token"
	"tokenalert_oauth-api/src/repositories"
	"github.com/rafawilliner/tokenalert_utils-go/src/rest_errors"
)

var (
	AccessTokenService accessTokenServiceInterface = &accessTokenService{}
)

type accessTokenService struct{}

type accessTokenServiceInterface interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestErr)	
}

func (a *accessTokenService) GetById(access_token_id string) (*access_token.AccessToken, rest_errors.RestErr) { 

	var at *access_token.AccessToken
	var err rest_errors.RestErr
	if at, err = repositories.AccessTokenRepository.GetById(access_token_id); err != nil {
		return nil, err
	}
	
	return at, nil
}
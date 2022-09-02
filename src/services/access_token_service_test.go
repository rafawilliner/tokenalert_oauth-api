package services

import (
	"errors"
	"testing"
	"tokenalert_oauth-api/src/domain/access_token"
	"tokenalert_oauth-api/src/repositories"

	"github.com/rafawilliner/tokenalert_utils-go/src/rest_errors"
	"github.com/stretchr/testify/assert"
)

var (
	getAccessTokenRepoFunc func(string) (*access_token.AccessToken, rest_errors.RestErr)
)

type accessTokenRepoMock struct{}

func (*accessTokenRepoMock) GetById(access_token_id string) (*access_token.AccessToken, rest_errors.RestErr) {
	return getAccessTokenRepoFunc(access_token_id)
}

func TestGetOK(t *testing.T) {

	accessToken := access_token.AccessToken{AccessToken: "ABC123", UserId: 2, ClientId: 3, Expires: 4 }
	getAccessTokenRepoFunc = func(access_token_id string) (*access_token.AccessToken, rest_errors.RestErr) {
		return &accessToken, nil
	}

	repositories.AccessTokenRepository = &accessTokenRepoMock{}

	_, err := AccessTokenService.GetById("ABC123")

	assert.NoError(t, err)
	assert.Equal(t, "ABC123", accessToken.AccessToken)
	assert.Equal(t, int64(4), accessToken.Expires)
}

func TestGetError(t *testing.T) {

	getAccessTokenRepoFunc = func(access_token_id string) (*access_token.AccessToken, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError("error fetching access token", errors.New("database error"))
	}

	repositories.AccessTokenRepository = &accessTokenRepoMock{}

	_, err := AccessTokenService.GetById("ABC123")

	assert.Error(t, err)
	assert.Equal(t, 500, err.Status())	
}

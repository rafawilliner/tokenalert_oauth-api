package services

import (
	"errors"
	"testing"
	"tokenalert_oauth-api/src/domain/access_token"
	"tokenalert_oauth-api/src/domain/users"
	"tokenalert_oauth-api/src/repositories/db"
	"tokenalert_oauth-api/src/repositories/rest"

	"github.com/rafawilliner/tokenalert_utils-go/src/rest_errors"
	"github.com/stretchr/testify/assert"
)

var (
	getAccessTokenRepoFunc func(string) (*access_token.AccessToken, rest_errors.RestErr)
	createAccessTokenRepoFunc func(access_token.AccessTokenRequest) rest_errors.RestErr
	loginUserRepoFunc func(string, string) (*users.User, rest_errors.RestErr)
)

type accessTokenRepoMock struct{}
type userRepoMock struct{}

func (*accessTokenRepoMock) GetById(access_token_id string) (*access_token.AccessToken, rest_errors.RestErr) {
	return getAccessTokenRepoFunc(access_token_id)
}

func (*accessTokenRepoMock) Create(atr access_token.AccessTokenRequest) rest_errors.RestErr {
	return createAccessTokenRepoFunc(atr)
}

func (*userRepoMock) LoginUser(email string, password string) (*users.User, rest_errors.RestErr) {
	return loginUserRepoFunc(email, password)
}

func TestGetOK(t *testing.T) {

	accessToken := access_token.AccessToken{AccessToken: "ABC123", UserId: 2, ClientId: 3, Expires: 4 }
	getAccessTokenRepoFunc = func(access_token_id string) (*access_token.AccessToken, rest_errors.RestErr) {
		return &accessToken, nil
	}

	db.AccessTokenRepository = &accessTokenRepoMock{}


	_, err := AccessTokenService.GetById("ABC123")

	assert.NoError(t, err)
	assert.Equal(t, "ABC123", accessToken.AccessToken)
	assert.Equal(t, int64(4), accessToken.Expires)
}

func TestGetByIdError(t *testing.T) {

	getAccessTokenRepoFunc = func(access_token_id string) (*access_token.AccessToken, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError("error fetching access token", errors.New("database error"))
	}

	db.AccessTokenRepository = &accessTokenRepoMock{}

	_, err := AccessTokenService.GetById("ABC123")

	assert.Error(t, err)
	assert.Equal(t, 500, err.Status())	
}

func TestCreateOK(t *testing.T) {

	accessTokenRequest := access_token.AccessTokenRequest{Email: "email@email.com", Password: "some-password", GrantType: "password" }
	user := users.User{ Id: 123, Name: "User1",  Email: "email@email.com", TelegramUser: "@user"  }
	createAccessTokenRepoFunc = func(atr access_token.AccessTokenRequest) rest_errors.RestErr {
		return nil
	}

	loginUserRepoFunc = func(email string, password string) (*users.User, rest_errors.RestErr) {
		return &user, nil
	}

	db.AccessTokenRepository = &accessTokenRepoMock{}
	rest.RestUsersRepository = &userRepoMock{}

	at, err := AccessTokenService.Create(accessTokenRequest)

	assert.NoError(t, err)
	assert.NotNil(t, at)
	assert.Equal(t, user.Id, at.UserId)
}

func TestCreateLoginUserFailError(t *testing.T) {

	accessTokenRequest := access_token.AccessTokenRequest{Email: "email@email.com", Password: "some-password", GrantType: "password" }
	createAccessTokenRepoFunc = func(atr access_token.AccessTokenRequest) rest_errors.RestErr {
		return nil
	}

	loginUserRepoFunc = func(email string, password string) (*users.User, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError("error login user in user api", errors.New("rest error"))
	}

	db.AccessTokenRepository = &accessTokenRepoMock{}
	rest.RestUsersRepository = &userRepoMock{}

	_, err := AccessTokenService.Create(accessTokenRequest)

	assert.Error(t, err)
	assert.Equal(t, 500, err.Status())	
	assert.Contains(t, "error login user in user api", err.Message())
}

func TestCreateValidateRequestFailed(t *testing.T) {

	accessTokenRequest := access_token.AccessTokenRequest{Email: "email@email.com", Password: "some-password" }
	createAccessTokenRepoFunc = func(atr access_token.AccessTokenRequest) rest_errors.RestErr {
		return nil
	}

	loginUserRepoFunc = func(email string, password string) (*users.User, rest_errors.RestErr) {
		return nil, nil
	}

	db.AccessTokenRepository = &accessTokenRepoMock{}
	rest.RestUsersRepository = &userRepoMock{}

	_, err := AccessTokenService.Create(accessTokenRequest)

	assert.Error(t, err)
	assert.Equal(t, 400, err.Status())	
	assert.Equal(t, "invalid grant_type parameter", err.Message())
}

func TestCreateTokenFailError(t *testing.T) {

	user := users.User{ Id: 123, Name: "User1",  Email: "email@email.com", TelegramUser: "@user"  }
	accessTokenRequest := access_token.AccessTokenRequest{Email: "email@email.com", Password: "some-password", GrantType: "password" }
	createAccessTokenRepoFunc = func(atr access_token.AccessTokenRequest) rest_errors.RestErr {
		return rest_errors.NewInternalServerError("error creating access token", errors.New("database error"))
	}

	loginUserRepoFunc = func(email string, password string) (*users.User, rest_errors.RestErr) {
		return &user, nil
	}

	db.AccessTokenRepository = &accessTokenRepoMock{}
	rest.RestUsersRepository = &userRepoMock{}

	_, err := AccessTokenService.Create(accessTokenRequest)

	assert.Error(t, err)
	assert.Equal(t, 500, err.Status())	
	assert.Contains(t, "error creating access token", err.Message())
}

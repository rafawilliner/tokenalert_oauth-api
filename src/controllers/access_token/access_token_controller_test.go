package access_token

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"tokenalert_oauth-api/src/domain/access_token"
	"tokenalert_oauth-api/src/services"

	"github.com/gin-gonic/gin"
	"github.com/rafawilliner/tokenalert_utils-go/src/rest_errors"
	"github.com/stretchr/testify/assert"
)

var (
	getAccessTokenFunc func(access_token_id string) (*access_token.AccessToken, rest_errors.RestErr)
	createAccessTokenFunc func(access_token.AccessTokenRequest) (*access_token.AccessToken, rest_errors.RestErr)
)

type accessTokenServiceMock struct{}

func (*accessTokenServiceMock) GetById(access_token_id string) (*access_token.AccessToken, rest_errors.RestErr) {
	return getAccessTokenFunc(access_token_id)
}

func (*accessTokenServiceMock) Create(at access_token.AccessTokenRequest) (*access_token.AccessToken, rest_errors.RestErr) {
	return createAccessTokenFunc(at)
}

func TestAccessTokenGetOK(t *testing.T) {

	getAccessTokenFunc = func(string) (*access_token.AccessToken, rest_errors.RestErr) {
		return &access_token.AccessToken{AccessToken: "123ABC", UserId: 1, ClientId: 1, Expires: 1}, nil
	}

	services.AccessTokenService = &accessTokenServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "/access_token", nil)
	c.Params = gin.Params{
		{Key: "access_token_id", Value: "123ABC"},
	}

	GetById(c)

	var accessTokenResponse access_token.AccessToken
	error := json.Unmarshal(response.Body.Bytes(), &accessTokenResponse)

	assert.Nil(t, error)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "123ABC", accessTokenResponse.AccessToken)
	assert.EqualValues(t, 1, accessTokenResponse.ClientId)
	assert.EqualValues(t, 1, accessTokenResponse.UserId)
	assert.EqualValues(t, 1, accessTokenResponse.Expires)
}

func TestUserGetBadRequestError(t *testing.T) {

	getAccessTokenFunc = func(access_token_id string) (*access_token.AccessToken, rest_errors.RestErr) {
		return nil, rest_errors.NewBadRequestError("wrong parameter format")
	}

	services.AccessTokenService = &accessTokenServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "/access_token", nil)
	c.Params = gin.Params{
		{Key: "access_token_bad_param", Value: "ABC"},
	}

	GetById(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)
}

func TestUserGetInternalServerError(t *testing.T) {

	getAccessTokenFunc = func(access_token_id string) (*access_token.AccessToken, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError("internal error", nil)
	}

	services.AccessTokenService = &accessTokenServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "/access_token", nil)
	c.Params = gin.Params{
		{Key: "access_token_id", Value: "ABC123"},
	}

	GetById(c)

	assert.EqualValues(t, http.StatusInternalServerError, response.Code)
}

func TestAccessTokenCreateOK(t *testing.T) {

	createAccessTokenFunc = func(atr access_token.AccessTokenRequest) (*access_token.AccessToken, rest_errors.RestErr) {
		return &access_token.AccessToken{AccessToken: "123ABC", UserId: 3, ClientId: 3, Expires: 3}, nil
	}

	services.AccessTokenService = &accessTokenServiceMock{}

	bodyAccessTokenRequest := map[string]interface{}{
		"AccessToken": "123ABC",
		"UserId": 3,
		"ClientId": 3,
		"Expires": 3,
	}

	body, _ := json.Marshal(bodyAccessTokenRequest)
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodPost, "/access_token", bytes.NewBuffer(body))	

	Create(c)

	var accessTokenResponse access_token.AccessToken
	error := json.Unmarshal(response.Body.Bytes(), &accessTokenResponse)

	assert.Nil(t, error)
	assert.EqualValues(t, http.StatusCreated, response.Code)
	assert.EqualValues(t, "123ABC", accessTokenResponse.AccessToken)
	assert.EqualValues(t, 3, accessTokenResponse.ClientId)
	assert.EqualValues(t, 3, accessTokenResponse.UserId)
	assert.EqualValues(t, 3, accessTokenResponse.Expires)
}

func TestAccessTokenCreateBadRequestError(t *testing.T) {

	createAccessTokenFunc = func(atr access_token.AccessTokenRequest) (*access_token.AccessToken, rest_errors.RestErr) {
		return nil, rest_errors.NewBadRequestError("wrong parameter format")
	}

	services.AccessTokenService = &accessTokenServiceMock{}

	bodyAccessTokenRequest := map[string]interface{}{
		"AccessToken": "123ABC",
		"UserId": 3,
		"ClientId": 3,
		"Expires": "ABC",
	}

	body, _ := json.Marshal(bodyAccessTokenRequest)
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodPost, "/access_token", bytes.NewBuffer(body))	

	Create(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)
}

func TestAccessTokenCreateInternalServerError(t *testing.T) {

	createAccessTokenFunc = func(atr access_token.AccessTokenRequest) (*access_token.AccessToken, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError("internal error", nil)
	}

	services.AccessTokenService = &accessTokenServiceMock{}

	bodyAccessTokenRequest := map[string]interface{}{
		"AccessToken": "123ABC",
		"UserId": 3,
		"ClientId": 3,
		"Expires": 3,
	}

	body, _ := json.Marshal(bodyAccessTokenRequest)
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodPost, "/access_token", bytes.NewBuffer(body))
	
	Create(c)

	assert.EqualValues(t, http.StatusInternalServerError, response.Code)
}

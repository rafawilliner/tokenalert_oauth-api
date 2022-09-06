package rest

import (
	"errors"
	"fmt"
	"os"
	"tokenalert_oauth-api/src/domain/users"
	"github.com/go-resty/resty/v2"
	"github.com/rafawilliner/tokenalert_utils-go/src/rest_errors"
)

var (
    Client *resty.Client
	RestUsersRepository restUsersRepositoryInterface = &restUsersRepository{}
) 

type restUsersRepository struct{}

type restUsersRepositoryInterface interface {	
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

func (r *restUsersRepository) LoginUser(email string, password string) (*users.User, rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	var userResponse *users.User
	var respError error

	//Client := resty.New()

	//print(fmt.Sprintf("%s/user/login", os.Getenv("user_api_url")))

	response, err := Client.R().
		SetBody(request).
		SetResult(&userResponse).
		SetError(&respError).
		Post(fmt.Sprintf("%s/user/login", os.Getenv("user_api_url")))

	if err != nil {
		panic(err)
	}

	if response.IsError() {
		switch response.StatusCode() {
		case 400:
			return nil, rest_errors.NewBadRequestError("invalid restclient response when trying to login user")

		case 404:
			return nil, rest_errors.NewNotFoundError("invalid restclient response when trying to login user")
		default:
			return nil, rest_errors.NewInternalServerError("invalid restclient response when trying to login user", errors.New("restclient error"))
		}		
	}

	return userResponse, nil
}

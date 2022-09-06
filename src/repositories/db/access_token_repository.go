package db

import (
	"context"
	"errors"
	"tokenalert_oauth-api/src/datasources/mongodb/access_token_db"
	"tokenalert_oauth-api/src/domain/access_token"

	"github.com/rafawilliner/tokenalert_utils-go/src/logger"
	"github.com/rafawilliner/tokenalert_utils-go/src/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	AccessTokenRepository accessTokenRepositoryInterface = &accessTokenRepository{}
)

type accessTokenRepository struct{}

type accessTokenRepositoryInterface interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestErr)
	Create(access_token.AccessTokenRequest) rest_errors.RestErr
}

func (a *accessTokenRepository) GetById(access_token_id string) (*access_token.AccessToken, rest_errors.RestErr) {

	var accessToken *access_token.AccessToken
	var accessTokenCollection *mongo.Collection = access_token_db.GetCollection("access_token")
	filter := bson.D{{Key: "access_token", Value: access_token_id}}
	err := accessTokenCollection.FindOne(context.TODO(), filter).Decode(&accessToken)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_errors.NewNotFoundError("document not found")
		} else {
			logger.Error("error when trying to get access token by id", err)
			return nil, rest_errors.NewInternalServerError("error fetching access token", errors.New("database error"))
		}
	}

	return accessToken, nil
}

func (a *accessTokenRepository) Create(atr access_token.AccessTokenRequest) rest_errors.RestErr {
	var accessTokenCollection *mongo.Collection = access_token_db.GetCollection("access_token")
	_, err := accessTokenCollection.InsertOne(context.TODO(), atr)
	if err != nil {
		logger.Error("error when trying to save new access token", err)
			return rest_errors.NewInternalServerError("error fetching access token", errors.New("database error"))
	}

	return nil
}

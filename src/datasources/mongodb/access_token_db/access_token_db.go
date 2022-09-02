package access_token_db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongo_user_name = "mongo_user_name"
	mongo_password = "mongo_password"
)

var (
	Db *mongo.Client
	username = os.Getenv(mongo_user_name)
	password = os.Getenv(mongo_password)
)

func InitDataBase() {
	var err error
    Db, err = mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.i1plsoa.mongodb.net/?retryWrites=true&w=majority", username, password)))
    if err != nil {
        panic(err)
    }
  
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = Db.Connect(ctx)
    if err != nil {
        panic(err)
    }

    err = Db.Ping(ctx, nil)
    if err != nil {
        panic(err)
    }
}

func GetCollection(collectionName string) *mongo.Collection {
    collection := Db.Database("oauth_db").Collection(collectionName)
    return collection
}

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
	DB *mongo.Client = ConnectDB()
	username = os.Getenv(mongo_user_name)
	password = os.Getenv(mongo_password)
)

func ConnectDB() *mongo.Client  {
    client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.i1plsoa.mongodb.net/?retryWrites=true&w=majority", username, password)))
    if err != nil {
        panic(err)
    }
  
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        panic(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        panic(err)
    }

    return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database("oauth_db").Collection(collectionName)
    return collection
}

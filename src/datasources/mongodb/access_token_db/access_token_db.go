package access_token_db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client = ConnectDB()

func ConnectDB() *mongo.Client  {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tokenalertapp:mongodbmamincho@cluster0.i1plsoa.mongodb.net/?retryWrites=true&w=majority"))
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

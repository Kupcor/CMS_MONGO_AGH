package api

import (
    "log"
    "fmt"
    "context"
	"historycznymonolog/constants"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    client *mongo.Client
)

func init() {
    clientOptions := options.Client().ApplyURI(constants.CONNECTION_STRING_TO_MONGODB)

    var err error
    client, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB...")
}
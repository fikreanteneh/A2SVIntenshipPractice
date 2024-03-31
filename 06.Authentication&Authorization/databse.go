
package main
import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var (
    mongoClient *mongo.Client
    once        sync.Once
)
// GetMongoClient returns a singleton instance of the MongoDB client.
func GetMongoClient() (*mongo.Client, error) {
    var err error
    once.Do(func() {
        uri := "mongodb://localhost:27017"
        clientOptions := options.Client().ApplyURI(uri)
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        mongoClient, err = mongo.Connect(ctx, clientOptions)
        if err != nil {
            log.Fatal(err)
        }

        err = mongoClient.Ping(ctx, nil)
        if err != nil {
            log.Println("Failed to connect to MongoDB")
            fmt.Println(err)
            return
        }

        fmt.Println("Successfully connected to MongoDB")
    })
    return mongoClient, err
}

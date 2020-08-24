package mongodb

import (
	"context"
	"log"
	"tkame123-net/worldcup-gq-server/adapter"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(hostURL string, database string) adapter.MongoClient {
	return func(ctx context.Context) (*adapter.MongoDBAPI, error) {
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(hostURL), options.Client().SetRetryWrites(false), options.Client().SetRetryReads(true))
		if err != nil {
			log.Fatalf("failed to open listener: %+v", err)
		}
		return &adapter.MongoDBAPI{
			MongoDBHostURL: hostURL,
			Database:       database,
			Client:         client,
		}, nil
	}
}

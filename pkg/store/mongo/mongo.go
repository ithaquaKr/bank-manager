package mongo

import (
	"context"

	"github.com/ithaquaKr/bank-manager/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongoClient(cfg *config.Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.Mongo.Uri)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	return client, nil
}

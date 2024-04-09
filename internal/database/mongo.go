package database

import (
	"context"
	"errors"
	"github.com/orewaee/embroidery-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

func GetCollection(collection string) *mongo.Collection {
	return database.Collection(collection)
}

func Load() error {
	uri := config.Get("MONGO_URI")
	if uri == "" {
		return errors.New("invalid mongo uri")
	}

	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	database = client.Database("embroidery")

	if err := client.Ping(context.Background(), nil); err != nil {
		return err
	}

	return nil
}

func Unload() error {
	return database.Client().Disconnect(context.Background())
}

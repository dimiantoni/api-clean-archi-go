package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongodbDatabase struct {
	Client *mongo.Client
}

func NewMongodb() *MongodbDatabase {

	dsn := os.Getenv("MONGODB_DSN")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	return &MongodbDatabase{Client: client}
}

func (m MongodbDatabase) GetConnection() any {
	return m.Client
}

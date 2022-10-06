package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() {
	clientOpt := options.Client().ApplyURI(os.Getenv("Connection"))

	client, err := mongo.Connect(context.TODO(), clientOpt)
	if err != nil {
		panic(err.Error())
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err.Error())
	}

	DB = client.Database(os.Getenv("DB"))
}

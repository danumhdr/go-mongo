package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() {
	clientOpt := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOpt)
	if err != nil {
		panic(err.Error())
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err.Error())
	}

	DB = client.Database("goDB")
}

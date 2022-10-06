package repository

import (
	"context"
	"go-mongo/config"
	"go-mongo/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

var collectionName string = "person"

func InsertPerson(data []interface{}) error {
	_, err := config.DB.Collection(collectionName).InsertMany(context.TODO(), data)
	return err
}

func UpdatePerson(filter interface{}, update interface{}) error {
	_, err := config.DB.Collection(collectionName).UpdateOne(context.TODO(), filter, update)
	return err
}

func GetPerson() []model.Person {
	var data []model.Person
	cur, err := config.DB.Collection(collectionName).Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Println(err.Error())
	}

	var response model.Person
	for cur.Next(context.Background()) {
		cur.Decode(&response)
		data = append(data, response)
	}

	return data
}

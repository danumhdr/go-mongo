package main

import (
	"go-mongo/config"
	"go-mongo/service"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	config.Connect()

	//command untuk insert
	// var person []interface{}
	// danu := model.Person{Name: "Danu", Age: 29, Id: "123"}
	// ika := model.Person{Name: "Ika", Age: 29, Id: "345"}
	// person = append(person, danu)
	// person = append(person, ika)
	// err := service.PersonInsert(person)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//command untuk update
	filter := bson.D{{Key: "name", Value: "Ika"}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "age", Value: 223},
		}},
	}

	e := service.PersonUpdate(filter, update)
	if e != nil {
		log.Println(e.Error())
	}

	//command untuk get
	data := service.PersonGet()
	log.Println(data)
}

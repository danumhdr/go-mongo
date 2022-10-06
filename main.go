package main

import (
	"go-mongo/config"
	"go-mongo/service"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

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
	filter := bson.D{{Key: "name", Value: "Danu"}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "age", Value: 99},
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

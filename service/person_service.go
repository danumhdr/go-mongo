package service

import (
	"go-mongo/model"
	"go-mongo/repository"
)

func PersonInsert(data []interface{}) error {
	getResult := repository.InsertPerson(data)
	return getResult
}

func PersonUpdate(filter interface{}, update interface{}) error {
	getResult := repository.UpdatePerson(filter, update)
	return getResult
}

func PersonGet() []model.Person {
	data := repository.GetPerson()
	return data
}

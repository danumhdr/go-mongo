package model

type Person struct {
	Name string `bson:"name"`
	Id   string `bson:"id"`
	Age  int    `bson:"age"`
}

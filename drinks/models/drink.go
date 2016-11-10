package models

import "gopkg.in/mgo.v2/bson"

//Drink represents a beverage
type Drink struct {
	ID    bson.ObjectId `bson:"_id"`
	Name  string        `json:"name"`
	Type1 string        `json:"type1"`
	Type2 string        `json:"type2"`
	Price string        `json:"price"`
	Stock string        `json:"stock"`
}

//Drinks is an array of Drink
type Drinks []Drink

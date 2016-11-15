package repository

import (
	"fmt"
	"log"

	"github.com/rest-go/drinks/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"	
)

// SERVER the DB server
const SERVER = "db:27017"

// DBNAME the name of the DB instance
const DBNAME = "dbdrinks"

// DOCNAME the name of the document
const DOCNAME = "drinks"

// GetDrinks returns the list of drinks
func GetDrinks() models.Drinks {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := models.Drinks{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// AddDrink inserts a drink in the DB
func AddDrink(drink models.Drink) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	drink.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(drink)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateDrink updates a drink in the DB
func UpdateDrink(drink models.Drink) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	drink.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).UpdateId(drink.ID, drink)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteDrink deletes a drink
func DeleteDrink(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "404"
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "500"
	}

	// Write status
	return "200"
}

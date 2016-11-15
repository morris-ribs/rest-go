# rest-go

This project consists of a REST API - developed in Go - that performs CRUD operations against a Mongo DB.

The concerned data are a list of drinks. Here is the description of a drink:

```Go
type Drink struct {
	ID    bson.ObjectId `bson:"_id"`
	Name  string        `json:"name"`
	Type1 string        `json:"type1"`
	Type2 string        `json:"type2"`
	Price string        `json:"price"`
	Stock string        `json:"stock"`
}
```

# Pre-requisites

1. Install and configure Go: https://golang.org/dl/
2. Install Mongo DB: https://docs.mongodb.com/manual/installation/?jmp=footer
3. Download Mgo project: `go get gopkg.in/mgo.v2`

# To run the project

## Start Mongo DB

1. In your terminal, type `mongod`
2. In another terminal, launch Mongo client: `mongo`
3. In the Mongo client, paste the contents of file mongo_insert.txt 

## Launch the project

In a terminal, go to the root folder of the project and launch it with the following command: `go run main.go`

## Running with Docker

To run the code with Docker:

1. Download and install Docker and docker-compose to your machine
2. Run the command `$ docker-compose up` in a terminal
3. Use an API client to make the calls

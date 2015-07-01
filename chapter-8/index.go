package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Category struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string
	Description string
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("taskdb").C("categories")
	c.RemoveAll(nil)
	// Index
	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	//insert three category objects
	err = c.Insert(
		&Category{bson.NewObjectId(), "Open-Source", "Tasks for open-source projects"},
		&Category{bson.NewObjectId(), "R & D", "R & D Tasks"},
		&Category{bson.NewObjectId(), "Project", "Project Tasks"},
	)
	if err != nil {
		panic(err)
	}

	result := Category{}
	err = c.Find(bson.M{"name": "Open-Source"}).One(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Description:", result.Description)
	}
}

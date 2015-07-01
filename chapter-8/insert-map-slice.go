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

	//get collection
	c := session.DB("taskdb").C("categories")
	c.RemoveAll(nil)
	doc := map[string]string{
		"name": "Open Source", "description": "Tasks for open-source projects",
	}

	//insert a map object
	err = c.Insert(doc)
	if err != nil {
		log.Fatal(err)
	}
	doc1 := bson.D{
		{"name", "Project"},
		{"description", "Project Tasks"},
	}
	err = c.Insert(doc1)
	if err != nil {
		log.Fatal(err)
	}
	var count int
	count, err = c.Count()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Count:", count)
}

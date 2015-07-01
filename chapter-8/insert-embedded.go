package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	Description string
	Due         time.Time
}
type Category struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string
	Description string
	Tasks       []Task
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
	doc := Category{
		bson.NewObjectId(),
		"Open Source",
		"Tasks for open-source projects",
		[]Task{
			Task{"Create project in mgo", time.Date(2015, time.August, 10, 0, 0, 0, 0, time.UTC)},
			Task{"Create REST API", time.Date(2015, time.August, 20, 0, 0, 0, 0, time.UTC)},
		},
	}
	//insert a Category object with embedded Tasks
	err = c.Insert(&doc)
	if err != nil {
		log.Fatal(err)
	}

	var count int
	count, err = c.Count()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d records inserted", count)

	iter := c.Find(nil).Iter()
	result := Category{}
	for iter.Next(&result) {
		fmt.Printf("Name:%s, Description:%s\n", result.Name, result.Description)
		tasks := result.Tasks
		for _, v := range tasks {
			fmt.Println(v.Description)
			fmt.Println(v.Due)
		}
	}
}

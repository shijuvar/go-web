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
	//Embedding child collection
	doc1 := Category{
		bson.NewObjectId(),
		"Open-Source",
		"Tasks for open-source projects",
		[]Task{
			Task{"Create project in mgo", time.Date(2015, time.August, 10, 0, 0, 0, 0, time.UTC)},
			Task{"Create REST API", time.Date(2015, time.August, 20, 0, 0, 0, 0, time.UTC)},
		},
	}
	doc2 := Category{
		bson.NewObjectId(),
		"GitHub",
		"Explore GitHub projects",
		[]Task{
			Task{"Evaluate Negroni Project", time.Date(2015, time.August, 15, 0, 0, 0, 0, time.UTC)},
			Task{"Explore mgo Project", time.Date(2015, time.August, 10, 0, 0, 0, 0, time.UTC)},
		},
	}
	//insert a Category object with embedded Tasks
	err = c.Insert(&doc1, &doc2)
	if err != nil {
		log.Fatal(err)
	}

	//get all records
	iter := c.Find(nil).Sort("name").Iter()
	result := Category{}
	for iter.Next(&result) {
		fmt.Printf("Category:%s, Description:%s\n", result.Name, result.Description)
		fmt.Println("-------------------------------------------")
		tasks := result.Tasks
		for _, v := range tasks {
			fmt.Printf("Task:%s Due:%v\n", v.Description, v.Due)
		}
		fmt.Println("-------------------------------------------")
	}
	if err = iter.Close(); err != nil {
		log.Fatal(err)
	}
	// get a single result
	result = Category{}
	err = c.Find(bson.M{"name": "Open-Source"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Category:%s, Description:%s\n", result.Name, result.Description)
	tasks := result.Tasks
	for _, v := range tasks {
		fmt.Printf("Task:%s Due:%v\n", v.Description, v.Due)
	}
	id := result.Id
	//updating the document
	err = c.Update(bson.M{"_id": id},
		bson.M{"$set": bson.M{
			"description": "Create open-source projects",
			"tasks": []Task{
				Task{"Evaluate Negroni Project", time.Date(2015, time.August, 15, 0, 0, 0, 0, time.UTC)},
				Task{"Explore mgo Project", time.Date(2015, time.August, 10, 0, 0, 0, 0, time.UTC)},
				Task{"Explore Gorilla Toolkit", time.Date(2015, time.August, 10, 0, 0, 0, 0, time.UTC)},
			},
		}})
	//Get the updated values
	result = Category{}
	err = c.FindId(id).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Category:%s, Description:%s\n", result.Name, result.Description)
	tasks = result.Tasks
	for _, v := range tasks {
		fmt.Printf("Task:%s Due:%v\n", v.Description, v.Due)
	}

}

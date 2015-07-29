package common

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
	}
	return session
}
func createDbSession() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
}

// Add indexes into MongoDB
func addIndexes() {
	var err error
	taskIndex := mgo.Index{
		Key:        []string{"createdby"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}
	noteIndex := mgo.Index{
		Key:        []string{"taskid"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}
	// Add indexes into MongoDB
	session := GetSession()
	defer session.Close()
	taskCol := session.DB("taskdb").C("tasks")
	noteCol := session.DB("taskdb").C("notes")
	err = taskCol.EnsureIndex(taskIndex)
	if err != nil {
		panic(err)
	}
	err = noteCol.EnsureIndex(noteIndex)
	if err != nil {
		panic(err)
	}
}

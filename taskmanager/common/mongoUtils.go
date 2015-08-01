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
	userIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
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
	session := GetSession().Copy()
	defer session.Close()
	userCol := session.DB("taskdb").C("users")
	taskCol := session.DB("taskdb").C("tasks")
	noteCol := session.DB("taskdb").C("notes")

	err = userCol.EnsureIndex(userIndex)
	if err != nil {
		panic(err)
	}
	err = taskCol.EnsureIndex(taskIndex)
	if err != nil {
		panic(err)
	}
	err = noteCol.EnsureIndex(noteIndex)
	if err != nil {
		panic(err)
	}
}

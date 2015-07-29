package controllers

import (
	"gopkg.in/mgo.v2"

	"github.com/shijuvar/go-web/taskmanager/common"
)

//Struct used for maintaining HTTP Request Context
type Context struct {
	MongoSession *mgo.Session
}

//Close mgo.Session
func (c *Context) Close() {
	c.MongoSession.Close()
}

//Returns a collection from the database.
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB("taskdb").C(name)
}

//Create a new Context object for each HTTP request
func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}

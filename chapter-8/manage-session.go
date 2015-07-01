package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

type (
	Category struct {
		Id          bson.ObjectId `bson:"_id,omitempty"`
		Name        string
		Description string
	}
	DataStore struct {
		session *mgo.Session
	}
)

//Close mgo.Session
func (d *DataStore) Close() {
	d.session.Close()
}

//Returns a collection from the database.
func (d *DataStore) C(name string) *mgo.Collection {
	return d.session.DB("taskdb").C(name)
}

//Create a new DataStore object for each HTTP request
func NewDataStore() *DataStore {
	ds := &DataStore{
		session: session.Copy(),
	}
	return ds
}

//Insert a record
func PostCategory(w http.ResponseWriter, r *http.Request) {
	var category Category
	// Decode the incoming Category json
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		panic(err)
	}
	ds := NewDataStore()
	defer ds.Close()
	//Getting the mgo.Collection
	c := ds.C("categories")
	//Insert record
	err = c.Insert(&category)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
}

//Read all records
func GetCategories(w http.ResponseWriter, r *http.Request) {

	var categories []Category
	ds := NewDataStore()
	defer ds.Close()
	//Getting the mgo.Collection
	c := ds.C("categories")
	iter := c.Find(nil).Iter()
	result := Category{}
	for iter.Next(&result) {
		categories = append(categories, result)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(categories)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func main() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/api/categories", GetCategories).Methods("GET")
	r.HandleFunc("/api/categories", PostCategory).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()

}

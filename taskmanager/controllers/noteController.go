package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shijuvar/go-web/taskmanager/data"
	"github.com/shijuvar/go-web/taskmanager/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var dataResource NoteResource
	// Decode the incoming Note json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		panic(err)
	}
	noteModel := dataResource.Data
	note := &models.TaskNote{
		TaskId:      bson.ObjectIdHex(noteModel.TaskId),
		Description: noteModel.Description,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("notes")
	//Insert a note document
	repo := &data.NoteRepository{c}
	repo.Create(note)
	if j, err := json.Marshal(note); err != nil {
		log.Fatal(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}
func GetNotesByTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("notes")
	repo := &data.NoteRepository{c}
	notes := repo.GetByTask(id)
	j, err := json.Marshal(NotesResource{Data: notes})
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
func GetNotes(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("notes")
	repo := &data.NoteRepository{c}
	notes := repo.GetAll()
	j, err := json.Marshal(NotesResource{Data: notes})
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
func GetNoteById(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("notes")
	repo := &data.NoteRepository{c}
	note, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	if j, err := json.Marshal(note); err != nil {
		log.Fatal(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource NoteResource
	// Decode the incoming Note json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		panic(err)
	}
	noteModel := dataResource.Data
	note := &models.TaskNote{
		Id:          id,
		Description: noteModel.Description,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("notes")
	repo := &data.NoteRepository{c}
	//Update note document
	if err := repo.Update(note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("notes")
	repo := &data.NoteRepository{c}
	//Delete a note document
	err := repo.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

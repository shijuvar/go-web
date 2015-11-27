package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shijuvar/go-web/taskmanager/common"
	"github.com/shijuvar/go-web/taskmanager/data"
	"github.com/shijuvar/go-web/taskmanager/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Handler for HTTP Post - "/notes"
// Insert a new Note document for a TaskId
func CreateNote(w http.ResponseWriter, r *http.Request) {
	var dataResource NoteResource
	// Decode the incoming Note json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Note data", 500)
		return
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
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

// Handler for HTTP Get - "/notes/tasks/{id}
// Returns all Notes documents under a TaskId
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
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Get - "/notes"
// Returns all Note documents
func GetNotes(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("notes")
	repo := &data.NoteRepository{c}
	notes := repo.GetAll()
	j, err := json.Marshal(NotesResource{Data: notes})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Get - "/notes/{id}"
// Returns a single Note document by id
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
			common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
			return
		}
	}
	if j, err := json.Marshal(note); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

// Handler for HTTP Put - "/notes/{id}"
// Update an existing Note document
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource NoteResource
	// Decode the incoming Note json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Note data", 500)
		return
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
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// Handler for HTTP Delete - "/notes/{id}"
// Delete an existing Note document
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
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

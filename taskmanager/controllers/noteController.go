package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/shijuvar/go-web/taskmanager/common"
	"github.com/shijuvar/go-web/taskmanager/data"
	"github.com/shijuvar/go-web/taskmanager/models"
)

// CreateNote inserts a new Note document for a TaskId
// Handler for HTTP Post - "/notes"
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
	col := context.DbCollection("notes")
	//Insert a note document
	repo := &data.NoteRepository{C: col}
	repo.Create(note)
	j, err := json.Marshal(note)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// GetNotesByTask returns all Notes documents under a TaskId
// Handler for HTTP Get - "/notes/tasks/{id}
func GetNotesByTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("notes")
	repo := &data.NoteRepository{C: col}
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

// GetNotes returns all Note documents
// Handler for HTTP Get - "/notes"
func GetNotes(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("notes")
	repo := &data.NoteRepository{C: col}
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

// GetNoteByID returns a single Note document by id
// Handler for HTTP Get - "/notes/{id}"
func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("notes")
	repo := &data.NoteRepository{C: col}
	note, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return

	}
	j, err := json.Marshal(note)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// UpdateNote updates an existing Note document
// Handler for HTTP Put - "/notes/{id}"
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
	col := context.DbCollection("notes")
	repo := &data.NoteRepository{C: col}
	//Update note document
	if err := repo.Update(note); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeleteNote deletes an existing Note document
// Handler for HTTP Delete - "/notes/{id}"
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("notes")
	repo := &data.NoteRepository{C: col}
	//Delete a note document
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

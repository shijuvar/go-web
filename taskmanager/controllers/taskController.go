package controllers

import (
	"encoding/json"
	"net/http"

	httpcontext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/shijuvar/go-web/taskmanager/common"
	"github.com/shijuvar/go-web/taskmanager/data"
)

// CreateTask insert a new Task document
// Handler for HTTP Post - "/tasks
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var dataResource TaskResource
	// Decode the incoming Task json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Task data",
			500,
		)
		return
	}
	task := &dataResource.Data
	context := NewContext()
	defer context.Close()
	if val, ok := httpcontext.GetOk(r, "user"); ok {
		context.User = val.(string)
	}
	task.CreatedBy = context.User
	col := context.DbCollection("tasks")
	repo := &data.TaskRepository{C: col}
	// Insert a task document
	repo.Create(task)
	j, err := json.Marshal(TaskResource{Data: *task})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)

}

// GetTasks returns all Task document
// Handler for HTTP Get - "/tasks"
func GetTasks(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tasks")
	repo := &data.TaskRepository{C: col}
	tasks := repo.GetAll()
	j, err := json.Marshal(TasksResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// GetTaskByID returns a single Task document by id
// Handler for HTTP Get - "/tasks/{id}"
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()

	defer context.Close()
	col := context.DbCollection("tasks")
	repo := &data.TaskRepository{C: col}
	task, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)

		} else {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)

		}
		return
	}
	j, err := json.Marshal(task)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// GetTasksByUser returns all Tasks created by a User
// Handler for HTTP Get - "/tasks/users/{id}"
func GetTasksByUser(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	user := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tasks")
	repo := &data.TaskRepository{C: col}
	tasks := repo.GetByUser(user)
	j, err := json.Marshal(TasksResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// UpdateTask update an existing Task document
// Handler for HTTP Put - "/tasks/{id}"
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource TaskResource
	// Decode the incoming Task json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Task data",
			500,
		)
		return
	}
	task := &dataResource.Data
	task.Id = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tasks")
	repo := &data.TaskRepository{C: col}
	// Update an existing Task document
	if err := repo.Update(task); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

// DeleteTask deelete an existing Task document
// Handler for HTTP Delete - "/tasks/{id}"
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tasks")
	repo := &data.TaskRepository{C: col}
	// Delete an existing Task document
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

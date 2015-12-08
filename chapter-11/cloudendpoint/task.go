package cloudendpoint

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

// Task is a datastore entity
type Task struct {
	Key         *datastore.Key `json:"id" datastore:"-"`
	Name        string         `json:"name" endpoints:"req"`
	Description string         `json:"description" datastore:",noindex" endpoints:"req"`
	CreatedOn   time.Time      `json:"createdon,omitempty"`
}

// Tasks is a response type of TaskService.List method
type Tasks struct {
	Tasks []Task `json:"tasks"`
}

// Struct is used to add API method
type TaskService struct {
}

// List returns a list of all the existing tasks.
func (ts *TaskService) List(c context.Context) (*Tasks, error) {
	tasks := []Task{}
	keys, err := datastore.NewQuery("tasks").Order("-CreatedOn").GetAll(c, &tasks)
	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		tasks[i].Key = k
	}
	return &Tasks{tasks}, nil
}

// Add inserts a new Task into Datastore
func (ts *TaskService) Add(c context.Context, t *Task) error {
	t.CreatedOn = time.Now()
	key := datastore.NewIncompleteKey(c, "tasks", nil)
	_, err := datastore.Put(c, key, t)
	return err
}

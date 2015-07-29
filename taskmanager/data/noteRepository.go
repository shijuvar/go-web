package data

import (
	"time"

	"github.com/shijuvar/go-web/taskmanager/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type NoteRepository struct {
	C *mgo.Collection
}

func (r *NoteRepository) Create(note *models.TaskNote) error {
	obj_id := bson.NewObjectId()
	note.Id = obj_id
	note.CreatedOn = time.Now()
	err := r.C.Insert(&note)
	return err
}

func (r *NoteRepository) Update(note *models.TaskNote) error {
	// partial update on MogoDB
	err := r.C.Update(bson.M{"_id": note.Id},
		bson.M{"$set": bson.M{
			"description": note.Description,
		}})
	return err
}
func (r *NoteRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
func (r *NoteRepository) GetByTask(id string) []models.TaskNote {
	var notes []models.TaskNote
	taskid := bson.ObjectIdHex(id)
	iter := r.C.Find(bson.M{"taskid": taskid}).Iter()
	result := models.TaskNote{}
	for iter.Next(&result) {
		notes = append(notes, result)
	}
	return notes
}
func (r *NoteRepository) GetAll() []models.TaskNote {
	var notes []models.TaskNote
	iter := r.C.Find(nil).Iter()
	result := models.TaskNote{}
	for iter.Next(&result) {
		notes = append(notes, result)
	}
	return notes
}
func (r *NoteRepository) GetById(id string) (note models.TaskNote, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&note)
	return
}

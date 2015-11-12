package task

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type Task struct {
	Name        string
	Description string
	CreatedOn   time.Time
}

const taskForm = `
<html>
  <body>
    <form action="/save" method="post">
      <p>Task Name: <input type="text" name="taskname" ></p>
	  <p> Description: <input type="text" name="description" ></p>
      <p><input type="submit" value="Submit"></p>
    </form>
  </body>
</html>
`
const taskListTmplHTML = `
<html>
<body>
<p>Task List</p>
{{range .}}
  <p>{{.Name}} - {{.Description}}</p>
{{end}}
<p><a href="/create">Create task</a> </p>
</body>
</html>
`

var taskListTemplate = template.Must(template.New("taskList").Parse(taskListTmplHTML))

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save", save)
}

func index(w http.ResponseWriter, r *http.Request) {
	q := datastore.NewQuery("tasks").
		Order("-CreatedOn")
	c := appengine.NewContext(r)	
	var tasks []Task
	_, err := q.GetAll(c, &tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := taskListTemplate.Execute(w, tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, taskForm)
}

func save(w http.ResponseWriter, r *http.Request) {
	task := Task{
		Name:        r.FormValue("taskname"),
		Description: r.FormValue("description"),
		CreatedOn:   time.Now(),
	}
	c := appengine.NewContext(r)
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "tasks", nil), &task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

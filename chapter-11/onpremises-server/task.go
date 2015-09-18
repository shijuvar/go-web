package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Task struct {
	Name        string
	Description string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/task", task)
	http.ListenAndServe(":8080",nil)
}

const taskForm = `
<html>
  <body>
    <form action="/task" method="post">
      <p>Task Name: <input type="text" name="taskname" ></p>
	  <p> Description: <input type="text" name="description" ></p>
      <p><input type="submit" value="Submit"></p>
    </form>
  </body>
</html>
`
const taskTemplateHTML = `
<html>
  <body>
    <p>New Task has been added:</p>
   <div>Task: {{.Name}}</div>
   <div>Description: {{.Description}}</div>
  </body>
</html>
`

var taskTemplate = template.Must(template.New("task").Parse(taskTemplateHTML))

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, taskForm)
}

func task(w http.ResponseWriter, r *http.Request) {
	task := Task{
		Name:        r.FormValue("taskname"),
		Description: r.FormValue("description"),
	}
	err := taskTemplate.Execute(w, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

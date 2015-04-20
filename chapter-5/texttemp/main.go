package main

import (
	"log"
	"os"
	"text/template"
)

type Note struct {
	Title       string
	Description string
}

const tmpl = `Note - Title: {{.Title}}, Description: {{.Description}}`

func main() {
	//define an instance of Note struct
	note := Note{"text/templates", "Template generates textual output"}

	//create a new template with a name
	t := template.New("note")

	//parse some content and generate a template
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	//Applies a parsed template to the data of Note object
	err1 := t.Execute(os.Stdout, note)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
		return
	}
}

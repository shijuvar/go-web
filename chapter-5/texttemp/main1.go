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

const tmpl = `Notes are:
{{range .}}
	 Title: {{.Title}}, Description: {{.Description}}
{{end}}
`

func main() {
	//define an instance
	notes := []Note{
		{"text/template", "Template generates textual output"},
		{"html/template", "Template generates HTML output"},
	}

	//create a new template with a name
	t := template.New("note")

	//parse some content and generate a template
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	//Applies a parsed template to the slice of Note objects
	err1 := t.Execute(os.Stdout, notes)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
		return
	}
}

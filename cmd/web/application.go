package web

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// tp for template.ExecuteTemplate
func tp(w http.ResponseWriter, d, n string, data interface{}) {
	fs, err := filepath.Glob(d)
	if err != nil {
		log.Fatalln(err)
	}
	tmpl := template.Must(template.ParseFiles(fs...))
	tmpl.ExecuteTemplate(w, n, data)
}

package web

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/GGCasa/luni"
)

var (
	d  = "./ui/html/ehut/*.gohtml"
	n  = "home"
	KV luni.Conn
)

type PathUrl struct {
	Url  string
	Nume string
}
type Test1 struct {
	Nume   string
	Numere []int
	Path   []PathUrl
}

// tp for template.ExecuteTemplate
func tp(w http.ResponseWriter, d, n string, data interface{}) {
	fs, err := filepath.Glob(d)
	if err != nil {
		log.Fatalln(err)
	}
	tmpl := template.Must(template.ParseFiles(fs...))
	tmpl.ExecuteTemplate(w, n, data)
}

func getTest1(k string) Test1 {
	var t1 Test1
	t1.Nume = "invalid"
	t1.Numere = []int{22, 33, 44, 55}
	t1.Path = []PathUrl{
		{Url: "/", Nume: "home link local"},
		{Url: "/test1", Nume: "test1 link valid"},
	}
	if k == "valid" {
		t1.Nume = "Nume test1"
		t1.Numere = []int{1, 2, 3}
		t1.Path = []PathUrl{
			{Url: "/", Nume: "home link valid"},
			{Url: "/test1", Nume: "test1 link local"},
		}
	}
	return t1
}

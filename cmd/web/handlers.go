package web

import (
	"net/http"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	data := getTest1("valid")

	tp(w, d, n, data)
}

func (a *app) test1(w http.ResponseWriter, r *http.Request) {
	data := getTest1("invalid")

	tp(w, d, n, data)
}

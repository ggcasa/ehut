package web

import (
	"net/http"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"t1": [2]int{1, 2},
		"t2": "Test 2",
		"t3": [3]string{"t1", "t2", "t3"},
		"t4": struct{ A string }{A: "struct t4"},
	}

	d := "./ui/html/ehut/*.gohtml"
	n := "home"

	tp(w, d, n, data)
}

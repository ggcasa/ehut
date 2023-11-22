package web

import (
	"fmt"
	"net/http"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Saluty")
}

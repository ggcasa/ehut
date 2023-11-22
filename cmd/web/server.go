package web

import (
	"log"
	"net/http"
)

type app struct{}

// Srv ehut web server
func Srv(p string) {
	a := app{}
	srv := http.Server{
		Addr:    p,
		Handler: a.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

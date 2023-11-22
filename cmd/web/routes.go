package web

import "net/http"

func (a *app) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", a.home)

	return mux
}

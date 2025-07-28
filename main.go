package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("This is the root path!"))
			return
		case "/home":
			w.Write([]byte("This is the home path!"))
			return
		default:
			w.Write([]byte("Not valid path!!!"))
			return
		}
	}
}

func main() {
	api := &api{addr: ":8880"}

	srv := &http.Server{
		Addr:    api.addr,
		Handler: api,
	}

	srv.ListenAndServe()
}

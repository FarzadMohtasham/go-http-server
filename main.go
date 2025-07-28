package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (s *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users list: ..."))
}

func (s *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a new user"))
}

func main() {
	api := &api{addr: ":8880"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	srv.ListenAndServe()
}

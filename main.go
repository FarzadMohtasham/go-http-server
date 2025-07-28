package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	s := &server{addr: ":8880"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(err)
	}
}

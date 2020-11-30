package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Storage interface {
}

// Server is a main application server and hancler structure
type Server struct {
	r   *mux.Router
	st  Storage
	srv *http.Server
}

// New creates a new server
func New(st Storage) *Server {
	srv := &Server{
		r:  mux.NewRouter(),
		st: st,
	}
	srv.srv = &http.Server{
		Handler: srv.r,
		Addr:    "localhost",
	}
	srv.setupRouter()
	return srv
}

func (s *Server) setupRouter() {
	s.r = mux.NewRouter()

	s.r.HandleFunc("/hello", s.hello).Methods("GET", "POST")
	s.r.HandleFunc("/", s.hello).Methods("GET", "POST")
	// s.r.HandleFunc("/user/{id}", s.getUser).Methods("GET")
	// s.r.HandleFunc("/user", s.setUser).Methods("POST")
	// s.r.HandleFunc("/user/{id}", s.setUser).Methods("PUT")
}

// Run server
func (s *Server) Run(addres string) error {
	srv := &http.Server{
		Handler: s.r,
		Addr:    addres,
	}

	return srv.ListenAndServe()
}

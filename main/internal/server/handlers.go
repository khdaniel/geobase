package server

import (
	// "encoding/json"
	"fmt"
	// "github.com/gorilla/mux"
	"net/http"
)

func (s *Server) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, request: %+v", r)
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id, ok := vars["id"]
	// if !ok {
	// 	http.Error(w, "parameter user is mandatory", http.StatusBadRequest)
	// 	return
	// }

}

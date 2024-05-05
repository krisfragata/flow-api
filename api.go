package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listeningAt string
}

func NewServer(listeningAt string) *Server {
	return &Server{
		listeningAt: listeningAt,
	}
}

func NotFoundError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 not found", http.StatusNotFound)
}

func (s *Server) Run(){
	router := mux.NewRouter()
	
	router.NotFoundHandler = http.HandlerFunc(NotFoundError)

	fmt.Println("Server running on port:", s.listeningAt)
	log.Fatal(http.ListenAndServe(s.listeningAt, router))

}



package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//HandleHelloWorld func writes string "Hello, Go!" into the ResponseWriter
func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, Go!")
	if err != nil {
		log.Printf("failed to write response %v", err)
	}
}

//HandleHelloWorld2 func writes string "Hello, Go!" into the ResponseWriter
func HandleHelloWorld2(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, Go! 2")
	if err != nil {
		log.Printf("failed to write response %v", err)
	}
}

func newRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleHelloWorld)
	mux.HandleFunc("/hello2", HandleHelloWorld2)
	return mux
}

func main() {
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      newRouter(),
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

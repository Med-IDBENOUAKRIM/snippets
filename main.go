package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatalln(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey There"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("snippet view"))
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("snippet form to create snippet"))
}

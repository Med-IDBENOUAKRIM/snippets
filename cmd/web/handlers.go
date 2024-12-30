package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
	}

	w.Header().Add("server", "Go")
	w.Write([]byte("Hey There"))
}

func getSnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// msg := fmt.Sprintf("Display a specific snippet with ID %d ... ", id)
	// w.Write([]byte(msg))
	fmt.Fprintf(w, "Display a specific snippet with ID %d ... ", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("snippet form to create snippet"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("snippet form to create snippet POST"))
}

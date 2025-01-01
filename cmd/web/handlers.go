package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/med-IDBENOUAKRIM/snippetbox/internal/models"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}

	w.Header().Add("server", "Go")
	w.Write([]byte("Hey There"))
}

func (app *Application) getSnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		// app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		return
	}

	// msg := fmt.Sprintf("Display a specific snippet with ID %d ... ", id)
	// w.Write([]byte(msg))
	fmt.Fprintf(w, "Display a specific snippet with ID %d ... ", id)
}

func (app *Application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("snippet form to create snippet"))
}

func (app *Application) createSnippet(w http.ResponseWriter, r *http.Request) {

	snippet := models.Snippet{
		Title:   "0 snail",
		Content: "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa",
		Expires: time.Now().Add(7 * 22 * time.Hour),
	}

	id, err := app.snippets.InsertSnippet(&snippet)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

}

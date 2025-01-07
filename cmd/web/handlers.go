package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/med-IDBENOUAKRIM/snippetbox/internal/models"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("server", "Go")
	snippets, err := app.snippets.LatestSnippets()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.render(w, r, http.StatusOK, "home.html", TemplateData{Snippets: snippets})
}

func (app *Application) getSnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		// app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		return
	}

	snippet, err := app.snippets.GetSnippet(id)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	app.render(w, r, http.StatusOK, "view.html", TemplateData{Snippet: snippet})

}

func (app *Application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("snippet form to create snippet"))
}

func (app *Application) createSnippet(w http.ResponseWriter, r *http.Request) {

	snippet := &models.Snippet{
		Title:   "2 snail",
		Content: "2 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa",
		Expires: time.Now().Add(7 * 22 * time.Hour),
	}

	err := app.snippets.InsertSnippet(snippet)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", snippet.ID), http.StatusSeeOther)

}

func (app *Application) getLastSnippet(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.snippets.LatestSnippets()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	fmt.Fprintf(w, "%+v", snippets)

}

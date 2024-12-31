package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) InsertSnippet(title, content string, expires int) (int, error) {
	return 0, nil
}

func (m *SnippetModel) GetSnippet() (Snippet, error) {
	return Snippet{}, nil
}

// get 10 most snippets recently created
func (m *SnippetModel) LatestSnippets() ([]Snippet, error) {
	return nil, nil
}

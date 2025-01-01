package models

import (
	"context"
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

// func (m *SnippetModel) InsertSnippet(title, content string, expires int) (int, error) {
func (m *SnippetModel) InsertSnippet(snippet *Snippet) (int, error) {
	query := `INSERT INTO snippets (title, content, expires) VALUES ($1, $2, $3) RETURNING id, title, created`

	args := []any{
		snippet.Title,
		snippet.Content,
		snippet.Expires,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) GetSnippet() (Snippet, error) {
	return Snippet{}, nil
}

// get 10 most snippets recently created
func (m *SnippetModel) LatestSnippets() ([]Snippet, error) {
	return nil, nil
}

/*
! ---> * --- DB.Query() is used for SELECT queries which return multiple rows.
! ---> * --- DB.QueryRow() is used for SELECT queries which return a single row.
! ---> * --- DB.Exec() is used for statements which don’t return rows (like INSERT and DELETE).
*/

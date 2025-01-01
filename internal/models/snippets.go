package models

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
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

func (m *SnippetModel) InsertSnippet(snippet *Snippet) error {
	query := `INSERT INTO snippets (title, content, expires) VALUES ($1, $2, $3) RETURNING id`

	args := []any{
		snippet.Title,
		snippet.Content,
		snippet.Expires,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&snippet.ID)
}

func (m *SnippetModel) GetSnippet(id int) (Snippet, error) {

	query := `SELECT id, title, content, created, expires FROM snippets WHERE expires > NOW() AND id = $1`

	var snippet Snippet

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrRecordNotFound
		} else {
			return Snippet{}, err
		}
	}

	return snippet, nil

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

package repository

import (
	"board/internal/handlers"
	"board/internal/repository/queries"
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type repo struct {
	*queries.Queries
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) Repository {
	return &repo{
		Queries: queries.NewQuery(conn),
		conn:    conn,
	}
}

type Repository interface {
	AddPost(context.Context, *handlers.Post) (int, error)
	GetPost(context.Context, *http.Request) (*handlers.Post, error)
	GetAllPosts(context.Context, *http.Request) ([]handlers.Post, error)
}

package handlers

import (
	"context"
	"net/http"
)

type boardRepo interface {
	AddPost(context.Context, *Post) (int, error)
	GetPost(context.Context, *http.Request) (*Post, error)
	GetAllPosts(context.Context, *http.Request) ([]Post, error)
}

type Post struct {
	Title string   `json:"title"`
	Price int      `json:"price"`
	Links []string `json:"photo"`
	Descr string   `json:"descr"`
}

type Brd struct {
	Board
	posts []Post
	repo  boardRepo
}

func NewPost() *Post {
	return &Post{
		Links: []string{},
	}
}

func NewBoard(repo boardRepo) Board {
	return &Brd{
		posts: make([]Post, 0),
		repo:  repo,
	}
}

type Board interface {
	AddPostHandler(w http.ResponseWriter, r *http.Request)
	GetAllPostsHandler(w http.ResponseWriter, r *http.Request)
	GetPostHandler(w http.ResponseWriter, r *http.Request)
}

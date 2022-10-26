package main

import (
	"board/internal/handlers"
	"board/internal/repository"
	"board/pkg/db"
	"context"
	"log"
	"net/http"
)

const dsn = `postgres://postgres:secret@db:5432/postgres`

func main() {
	conn, err := db.New(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	repo := repository.NewRepository(conn)
	b := handlers.NewBoard(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/get_posts", b.GetAllPostsHandler)
	mux.HandleFunc("/get_post", b.GetPostHandler)
	mux.HandleFunc("/add_post", b.AddPostHandler)
	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}

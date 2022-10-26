package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (b *Brd) AddPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer r.Body.Close()

	buf := NewPost()
	if err := json.Unmarshal(body, buf); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	id, err := b.repo.AddPost(context.Background(), buf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("PostID:%d\n", id)))
}

func (b *Brd) GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	posts, err := b.repo.GetAllPosts(context.Background(), r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(posts)
}

func (b *Brd) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	if _, ok := r.URL.Query()["id"]; ok {
		w.Header().Set("Content-Type", "application/json")
		post, err := b.repo.GetPost(context.Background(), r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No ID in request"))
	}
}

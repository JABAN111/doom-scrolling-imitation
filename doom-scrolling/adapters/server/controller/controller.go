package controller

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"rshd/lab1/v2/core"
)

func NewCreateUserHandler(ctx context.Context, log *slog.Logger, service *core.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user core.User
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request payload"))
			return
		}

		if user.Username == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Username is required"))
			return
		}

		_, err = service.CreateUser(ctx, user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to create user"))
			return
		}

		w.Write([]byte("User created"))
		w.WriteHeader(http.StatusCreated)
	}
}

func NewCreatePostHandler(ctx context.Context, log *slog.Logger, service *core.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post core.Post
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&post)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request payload"))
			return
		}

		if post.ID == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ID is required"))
			return
		}

		_, err = service.CreatePost(ctx, post)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to create post"))
			return
		}

		w.Write([]byte("Post created"))
		w.WriteHeader(http.StatusCreated)
	}
}

func NewFollowUserHandler(ctx context.Context, log *slog.Logger, service *core.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := service.FollowUser(ctx, r.URL.Query().Get("username"), r.URL.Query().Get("usernameToFollow"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to follow user"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User followed"))
	}
}

func NewLikeHandler(ctx context.Context, log *slog.Logger, service *core.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := service.LikePost(ctx, r.URL.Query().Get("user_id"), r.URL.Query().Get("post_id"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to like post"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Post liked"))
	}
}

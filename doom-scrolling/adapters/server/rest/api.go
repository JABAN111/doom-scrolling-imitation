package rest

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"rshd/lab1/v2/core"
	"rshd/lab1/v2/util"
)

func NewCreateUserHandler(ctx context.Context, log *slog.Logger, service *core.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user core.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			util.WriteResponse(ctx, log, w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		if user.Username == "" {
			util.WriteResponse(ctx, log, w, http.StatusBadRequest, "Username is required")
			return
		}
		_, err := service.CreateUser(ctx, user)
		if err != nil {
			util.WriteResponse(ctx, log, w, http.StatusInternalServerError, "Failed to create user")
			return
		}
		util.WriteResponse(ctx, log, w, http.StatusCreated, "User created")
	}
}

func NewCreatePostHandler(ctx context.Context, log *slog.Logger, service *core.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post core.Post
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			util.WriteResponse(ctx, log, w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		if post.ID == "" {
			util.WriteResponse(ctx, log, w, http.StatusBadRequest, "ID is required")
			return
		}
		_, err := service.CreatePost(ctx, post)
		if err != nil {
			util.WriteResponse(ctx, log, w, http.StatusInternalServerError, "Failed to create post")
			return
		}
		util.WriteResponse(ctx, log, w, http.StatusCreated, "Post created")
	}
}

func NewFollowUserHandler(ctx context.Context, log *slog.Logger, service *core.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := service.FollowUser(ctx, r.URL.Query().Get("username"), r.URL.Query().Get("usernameToFollow"))
		if err != nil {
			util.WriteResponse(ctx, log, w, http.StatusInternalServerError, "Failed to follow user")
			return
		}
		util.WriteResponse(ctx, log, w, http.StatusOK, "User followed")
	}
}

func NewLikeHandler(ctx context.Context, log *slog.Logger, service *core.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := service.LikePost(ctx, r.URL.Query().Get("user_id"), r.URL.Query().Get("post_id"))
		if err != nil {
			util.WriteResponse(ctx, log, w, http.StatusInternalServerError, "Failed to like post")
			return
		}
		util.WriteResponse(ctx, log, w, http.StatusOK, "Post liked")
	}
}

func NewFeed(log *slog.Logger, service *core.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := service.GetFeed(r.Context(), r.URL.Query().Get("username"))
		if err != nil {
			util.WriteResponse(r.Context(), log, w, http.StatusInternalServerError, "Failed to get feed")
			return
		}
		util.WriteResponseJSON(r.Context(), log, w, http.StatusOK, res)
	}
}

package rest

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"rshd/lab1/v2/core"
	"rshd/lab1/v2/util"
	"strconv"
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

func NewFeedHandler(log *slog.Logger, service *core.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := service.GetFeed(r.Context(), r.URL.Query().Get("username"))
		if err != nil {
			util.WriteResponse(r.Context(), log, w, http.StatusInternalServerError, "Failed to get feed")
			return
		}
		util.WriteResponseJSON(r.Context(), log, w, http.StatusOK, res)
	}
}

func NewUserStats(log *slog.Logger, service *core.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		if username == "" {
			util.WriteResponse(r.Context(), log, w, http.StatusBadRequest, "Username is required")
			return
		}

		res, err := service.GetUserStats(r.Context(), username)
		if err != nil {
			util.WriteResponse(r.Context(), log, w, http.StatusInternalServerError, "Failed to get stats")
			return
		}
		util.WriteResponseJSON(r.Context(), log, w, http.StatusOK, res)
	}
}

func NewPopularTags(log *slog.Logger, service *core.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		days := r.URL.Query().Get("days")
		limit := r.URL.Query().Get("limit")

		if days == "" || limit == "" {
			util.WriteResponse(r.Context(), log, w, http.StatusBadRequest, "days & limit is required")
			return
		}
		daysInt, err := strconv.Atoi(days)
		if err != nil {
			util.WriteResponse(r.Context(), log, w, http.StatusBadRequest, "days must be an integer")
			return
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			util.WriteResponse(r.Context(), log, w, http.StatusBadRequest, "limit must be an integer")
			return
		}

		tags, err := service.GetPopularTags(r.Context(), daysInt, limitInt)
		if err != nil {
			log.Error("Failed to get popular tags", "error", err)
			util.WriteResponse(r.Context(), log, w, http.StatusInternalServerError, "Failed to get tags")
			return
		}
		util.WriteResponseJSON(r.Context(), log, w, http.StatusOK, tags)
	}
}

package rest

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"rshd/lab1/v2/core"
	"rshd/lab1/v2/core/service"
	"rshd/lab1/v2/internal/util"
	"strconv"
)

func NewCreateUserHandler(ctx context.Context, log *slog.Logger, service *service.Service) http.HandlerFunc {
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

func NewCreatePostHandler(ctx context.Context, log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			util.WriteResponse(ctx, log, w, http.StatusBadRequest, "Invalid multipart form")
			return
		}

		postData := r.FormValue("post")
		var post core.Post
		if err := json.Unmarshal([]byte(postData), &post); err != nil {
			util.WriteResponse(ctx, log, w, http.StatusBadRequest, "Invalid post JSON")
			return
		}
		if post.ID == "" {
			util.WriteResponse(ctx, log, w, http.StatusBadRequest, "ID is required")
			return
		}
		file, header, err := r.FormFile("image")
		if err != nil {
			util.WriteResponse(ctx, log, w, http.StatusBadRequest, "Image file is required")
			return
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				return
			}
		}(file)
		tmp, err := os.CreateTemp("", "upload-*"+filepath.Ext(header.Filename))
		if err != nil {
			util.WriteResponse(ctx, log, w, http.StatusInternalServerError, "Cannot create temp file")
			return
		}
		defer func() {
			err := tmp.Close()
			if err != nil {
				return
			}
			err = os.Remove(tmp.Name())
			if err != nil {
				return
			}
		}()

		if _, err := io.Copy(tmp, file); err != nil {
			util.WriteResponse(ctx, log, w, http.StatusInternalServerError, "Failed to save uploaded image")
			return
		}

		_, err = service.CreatePost(ctx, post, tmp.Name())
		if err != nil {
			util.WriteResponse(ctx, log, w, http.StatusInternalServerError, "Failed to create post")
			return
		}
		util.WriteResponse(ctx, log, w, http.StatusCreated, "Post created")
	}
}

func NewFollowUserHandler(ctx context.Context, log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := service.FollowUser(ctx, r.URL.Query().Get("username"), r.URL.Query().Get("usernameToFollow"))
		if err != nil {
			util.WriteResponse(ctx, log, w, http.StatusInternalServerError, "Failed to follow user")
			return
		}
		util.WriteResponse(ctx, log, w, http.StatusOK, "User followed")
	}
}

func NewLikeHandler(ctx context.Context, log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := service.LikePost(ctx, r.URL.Query().Get("user_id"), r.URL.Query().Get("post_id"))
		if err != nil {
			util.WriteResponse(ctx, log, w, http.StatusInternalServerError, "Failed to like post")
			return
		}
		util.WriteResponse(ctx, log, w, http.StatusOK, "Post liked")
	}
}

func NewFeedHandler(log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := service.GetFeed(r.Context(), r.URL.Query().Get("username"))
		if err != nil {
			util.WriteResponse(r.Context(), log, w, http.StatusInternalServerError, "Failed to get feed")
			return
		}
		util.WriteResponseJSON(r.Context(), log, w, http.StatusOK, res)
	}
}

func NewUserStats(log *slog.Logger, service *service.Service) http.HandlerFunc {
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

func NewPopularTags(log *slog.Logger, service *service.Service) http.HandlerFunc {
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

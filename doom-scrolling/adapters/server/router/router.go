package router

import (
	"context"
	"log/slog"
	"net/http"
	"rshd/lab1/v2/adapters/server/controller"
	"rshd/lab1/v2/core"
)

func RegisterRoutes(ctx context.Context, log *slog.Logger, mux *http.ServeMux, service *core.Service) {
	mux.HandleFunc("POST /api/create", controller.NewCreateUserHandler(
		ctx,
		log,
		service,
	))

	mux.HandleFunc("POST /api/create/post", controller.NewCreatePostHandler(ctx, log, service))
	mux.HandleFunc("POST /api/follow", controller.NewFollowUserHandler(ctx, log, service))
	mux.HandleFunc("POST /api/like", controller.NewLikeHandler(ctx, log, service))
	mux.HandleFunc("GET /api/feed", controller.NewFeed(log, service))
}

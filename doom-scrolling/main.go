package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"rshd/lab1/v2/adapters/db/couchbase"
	"rshd/lab1/v2/adapters/db/neof4j"
	"rshd/lab1/v2/config"
	"rshd/lab1/v2/core"
	"time"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg := config.Config{
		CouchBaseCfg: config.CouchBaseConfig{
			URL:      "localhost",
			Username: "jaba_admin",
			Password: "jaba_pwd",
			Bucket:   "doom-scrolling",
		},
		DgraphCfg: config.DgraphConfig{
			URL: "localhost:9080",
		},
	}

	docDB, err := couchbase.New(logger, cfg)
	if err != nil {
		logger.Error("Failed to initialize Couchbase", "error", err)
		os.Exit(1)
	}

	graphDB, err := neof4j.New(logger, cfg)
	if err != nil {
		logger.Error("Failed to initialize Neo4j", "error", err)
		os.Exit(1)
	}
	// defer graphDB.Close()

	service := core.NewService(logger, docDB, graphDB)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	users := []string{"alice", "bob", "charlie"}
	for _, username := range users {

		_, err := service.CreateUser(ctx, core.User{Username: username})
		if err != nil {
			logger.Error("Failed to create user", "username", username, "error", err)
		} else {
			fmt.Printf("User %s created\n", username)
		}
	}

	if err := service.FollowUser(ctx, "alice", "bob"); err != nil {
		logger.Error("alice follow bob failed", "error", err)
	}
	if err := service.FollowUser(ctx, "alice", "charlie"); err != nil {
		logger.Error("alice follow charlie failed", "error", err)
	}

	if err := service.FollowUser(ctx, "bob", "charlie"); err != nil {
		logger.Error("bob follow charlie failed", "error", err)
	}

	if _, err := service.CreatePost(ctx, core.Post{ID: "post1", UserID: "bob"}); err != nil {
		logger.Error("bob post1 creation failed", "error", err)
	}
	time.Sleep(100 * time.Millisecond)
	if _, err := service.CreatePost(ctx, core.Post{ID: "post2", UserID: "bob"}); err != nil {
		logger.Error("bob post2 creation failed", "error", err)
	}

	time.Sleep(100 * time.Millisecond)
	if _, err := service.CreatePost(ctx, core.Post{ID: "post3", UserID: "charlie"}); err != nil {
		logger.Error("charlie post3 creation failed", "error", err)
	}
	time.Sleep(100 * time.Millisecond)
	if _, err := service.CreatePost(ctx, core.Post{ID: "post4", UserID: "charlie"}); err != nil {
		logger.Error("charlie post4 creation failed", "error", err)
	}

	if err := service.LikePost(ctx, "alice", "post1"); err != nil {
		logger.Error("alice like post1 failed", "error", err)
	}

	aliceFeed, err := service.GetFeed(ctx, "alice")
	if err != nil {
		logger.Error("Get feed for alice failed", "error", err)
	} else {
		fmt.Printf("Alice's feed: %v\n", aliceFeed)

	}

	bobFeed, err := service.GetFeed(ctx, "bob")
	if err != nil {
		logger.Error("Get feed for bob failed", "error", err)
	} else {
		fmt.Printf("Bob's feed: %v\n", bobFeed)

	}

	charlieFeed, err := service.GetFeed(ctx, "charlie")
	if err != nil {
		logger.Error("Get feed for charlie failed", "error", err)
	} else {
		fmt.Printf("Charlie's feed: %v\n", charlieFeed)
	}
}

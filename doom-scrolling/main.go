package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"rshd/lab1/v2/adapters/db/clickhouse"
	"rshd/lab1/v2/adapters/db/couchbase"
	"rshd/lab1/v2/adapters/db/influx"
	"rshd/lab1/v2/adapters/db/neof4j"
	"rshd/lab1/v2/adapters/server/rest"
	"rshd/lab1/v2/adapters/sss"
	"rshd/lab1/v2/config"
	"rshd/lab1/v2/core"
	"rshd/lab1/v2/core/service"
	"time"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	crg := config.MustLoad("./config.yaml")
	fmt.Printf("data %v \n", crg)

	endpoint := crg.Minio.URL

	// Initialize minio client object.
	minioClient, err := sss.NewMinio(log, endpoint, false)
	if err != nil {
		if errors.Is(err, sss.ErrAccessKeyId) {
			log.Error("access key env is not specified", "err", err)
			os.Exit(1)
		}
		if errors.Is(err, sss.ErrSecretAccessKey) {
			log.Error("secret key env is not specified", "err", err)
			os.Exit(1)
		}
		log.Error("unexpected error while connecting to minio", "err", err)
		os.Exit(2)
	}

	testFile := "./config.yaml"
	fmt.Println(filepath.Base(testFile))
	err = minioClient.UploadPostImage(context.Background(), filepath.Base(testFile), testFile)
	if err != nil {
		panic(err)
	}
	log.Info("minio working")

	docDB, err := couchbase.New(log, crg)
	if err != nil {
		log.Error("Failed to initialize Couchbase", "error", err)
		os.Exit(1)
	}

	graphDB, err := neof4j.New(log, crg)
	if err != nil {
		log.Error("Failed to initialize Neo4j", "error", err)
		os.Exit(1)
	}

	token := "0YVmO2e179ymcr4AZoA9FOEAIZSdDmezA8yIuLnSL4ERowgKZGKWEKqZAR64BCVn1aC4tN6Jq7aVM0ldAMZJIQ=="
	influxClient, err := influx.New(log, "http://"+crg.Influx.URL, token, "docs", "home")
	if err != nil {
		log.Error("Failed to initialize influx db", "error", err)
		os.Exit(1)
	}

	clickhouseClient := clickhouse.New(log, crg.ClickhouseConfig.URL)

	numWorkers := 20
	service := service.NewService(log, docDB, graphDB, influxClient, clickhouseClient, minioClient, numWorkers)
	fmt.Print(service)

	ctx := context.Background()

	users := []string{"alice", "bob", "charlie"}
	for _, username := range users {
		_, err := service.CreateUser(ctx, core.User{
			Username: username,
			Email:    username + "@gmail.com",
			Bio:      username + username + username,
			Age:      22,
		})
		if err != nil {
			log.Error("Failed to create user", "username", username, "error", err)
		} else {
			fmt.Printf("User %s created\n", username)
		}
	}

	if err := service.FollowUser(ctx, "alice", "bob"); err != nil {
		log.Error("alice follow bob failed", "error", err)
	}
	if err := service.FollowUser(ctx, "alice", "charlie"); err != nil {
		log.Error("alice follow charlie failed", "error", err)
	}

	if err := service.FollowUser(ctx, "bob", "charlie"); err != nil {
		log.Error("bob follow charlie failed", "error", err)
	}

	if _, err := service.CreatePost(ctx, core.Post{ID: "post1", UserID: "bob"}, testFile); err != nil {
		log.Error("bob post1 creation failed", "error", err)
	}
	time.Sleep(100 * time.Millisecond)
	if _, err := service.CreatePost(ctx, core.Post{ID: "post2", UserID: "bob"}, testFile); err != nil {
		log.Error("bob post2 creation failed", "error", err)
	}

	time.Sleep(100 * time.Millisecond)
	if _, err := service.CreatePost(ctx, core.Post{ID: "post3", UserID: "charlie"}, testFile); err != nil {
		log.Error("charlie post3 creation failed", "error", err)
	}
	time.Sleep(100 * time.Millisecond)
	if _, err := service.CreatePost(ctx, core.Post{ID: "post4", UserID: "charlie"}, testFile); err != nil {
		log.Error("charlie post4 creation failed", "error", err)
	}

	if err := service.LikePost(ctx, "alice", "post1"); err != nil {
		log.Error("alice like post1 failed", "error", err)
	}

	aliceFeed, err := service.GetFeed(ctx, "alice")
	if err != nil {
		log.Error("Get feed for alice failed", "error", err)
	} else {
		fmt.Printf("Alice's feed: %v\n", aliceFeed)

	}

	bobFeed, err := service.GetFeed(ctx, "bob")
	if err != nil {
		log.Error("Get feed for bob failed", "error", err)
	} else {
		fmt.Printf("Bob's feed: %v\n", bobFeed)

	}

	charlieFeed, err := service.GetFeed(ctx, "charlie")
	if err != nil {
		log.Error("Get feed for charlie failed", "error", err)
	} else {
		fmt.Printf("Charlie's feed: %v\n", charlieFeed)
	}

	if err != nil {
		log.Error("Ошибка вставки данных: %v", "err", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/action/create", rest.NewCreateUserHandler(
		ctx,
		log,
		service,
	))
	mux.HandleFunc("POST /api/action/create/post", rest.NewCreatePostHandler(ctx, log, service))

	mux.HandleFunc("POST /api/action/follow", rest.NewFollowUserHandler(ctx, log, service))
	mux.HandleFunc("POST /api/action/like", rest.NewLikeHandler(ctx, log, service))
	mux.HandleFunc("GET /api/action/feed", rest.NewFeedHandler(log, service))

	mux.HandleFunc("GET /api/stat", rest.NewUserStats(log, service))
	mux.HandleFunc("GET /api/stat/popular", rest.NewPopularTags(log, service))
	res, err := service.GetUserStats(ctx, "charlie")
	if err != nil {
		log.Error("Failed to get user stats")
	}
	log.Info("res", "res", res)

	na, err := service.GetPopularTags(ctx, 10, 10)
	if err != nil {
		log.Error("Failed to get popular tags")
	}
	for _, r := range na {
		fmt.Println(r)
	}

	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      mux,
	}

	go service.StartSystemMetricsCollection(context.Background(), time.Minute*1)
	log.Info("server is starting...")
	if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Error("server closed unexpectedly", "error", err)
	}
}

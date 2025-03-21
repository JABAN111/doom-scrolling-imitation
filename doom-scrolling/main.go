package main

import (
	"context"
	"rshd/lab1/v2/adapters/db/couchbase"
	"rshd/lab1/v2/adapters/db/dgraph"
	"rshd/lab1/v2/config"
	"rshd/lab1/v2/core"
	"rshd/lab1/v2/logger"
)

func main() {
	log := logger.GetInstance()
	ctx := context.Background()

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
	couchBase, err := couchbase.New(log, cfg)
	if err != nil {
		log.Error("Failed to create couchbase", "error", err)
		panic(err)
	}
	dgraph, err := dgraph.New(log, cfg)

	// dgraph.RANDOMBULLSHITGO()
	// dgraph.READ()
	// dgraph.SETBEN200()
	// os.Exit(0)

	if err != nil {
		log.Error("Failed to create dgraph", "error", err)
		panic(err)
	}
	log.Info("Successfully connected to Couchbase and Dgraph")
	service := core.NewService(log, couchBase, dgraph)
	user1, err := service.CreateUser(ctx, core.User{Username: "john_doe", Email: "jaba@yadro.com", Bio: "Photographer"})
	user2, err := service.CreateUser(ctx, core.User{Username: "jaba", Email: "jaba@yadro.com", Bio: "Photographer"})

	if err != nil {
		panic(err)
	}
	log.Info("User created")

	post1 := core.Post{ID: "post_789", UserID: "user_456", ImageURL: "https://cdn.example.com/photo.jpg", Caption: "Sunset in Bali", Likes: 0, CreatedAt: "2024-02-10T10:15:00Z"}
	_, err = service.CreatePost(ctx, post1)
	if err != nil {
		panic(err)
	}
	log.Info("Post created")
	err = service.LikePost(ctx, user1.Username, post1.ID) //Не работает???
	if err != nil {
		panic(err)
	}

	err = service.FollowUser(ctx, user1.Username, user2.Username)
	if err != nil {
		panic(err)
	}
	log.Info("User followed")

	// res, err := service.GetFeed(ctx, user1.Username)
	if err != nil {
		panic(err)
	}
	// log.Info("Feed received", "feed", res)
}

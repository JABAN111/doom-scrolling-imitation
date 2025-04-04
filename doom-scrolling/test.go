package main

// import (
// 	"context"
// 	"log"
// 	"log/slog"
// 	"os"
// 	"rshd/lab1/v2/adapters/db/neof4j"
// 	"rshd/lab1/v2/config"
// 	"time"
// )

// func main() {
// 	// Инициализация логгера
// 	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
// 	cfg := config.Config{
// 		CouchBaseCfg: config.CouchBaseConfig{
// 			URL:      "localhost",
// 			Username: "jaba_admin",
// 			Password: "jaba_pwd",
// 			Bucket:   "doom-scrolling",
// 		},
// 		DgraphCfg: config.DgraphConfig{
// 			URL: "localhost:9080",
// 		},
// 	}
// 	// Конфигурация (в реальном приложении берется из env/config файла)

// 	// Инициализация подключения
// 	db, err := neof4j.New(logger, cfg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// Создание пользователей
// 	if err := db.CreateUser(ctx, "alice"); err != nil {
// 		log.Fatal("Failed to create user alice:", err)
// 	}

// 	if err := db.CreateUser(ctx, "bob"); err != nil {
// 		log.Fatal("Failed to create user bob:", err)
// 	}

// 	// Alice подписывается на Bob
// 	if err := db.FollowUser(ctx, "alice", "bob"); err != nil {
// 		log.Fatal("Follow failed:", err)
// 	}

// 	// Bob создает пост
// 	if err := db.CreatePost(ctx, "post1", "bob"); err != nil {
// 		log.Fatal("Create post failed:", err)
// 	}

// 	// Alice лайкает пост Bob'а
// 	if err := db.LikePost(ctx, "alice", "post1"); err != nil {
// 		log.Fatal("Like failed:", err)
// 	}

// 	// Получаем ленту Alice
// 	feed, err := db.GetFeed(ctx, "alice")
// 	if err != nil {
// 		log.Fatal("Get feed failed:", err)
// 	}

// 	log.Println("Alice's feed:", feed) // Должен содержать ["post1"]
// }

package main

import (
	"context"
	"net/http"
	"rshd/lab1/v2/adapters/db/couchbase"
	"rshd/lab1/v2/adapters/db/dgraph"
	"rshd/lab1/v2/adapters/server"
	"rshd/lab1/v2/adapters/server/router"
	"rshd/lab1/v2/config"
	"rshd/lab1/v2/core"
	"rshd/lab1/v2/logger"
	"time"
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
	if err != nil {
		log.Error("Failed to create dgraph", "error", err)
		panic(err)
	}

	log.Info("Successfully connected to Couchbase and Dgraph")
	service := core.NewService(log, couchBase, dgraph)

	rootMux := http.NewServeMux()
	router.RegisterRoutes(ctx, log, rootMux, service)

	restServer := server.NewServer(log, rootMux, "localhost:8082", time.Second*10)
	restServer.Run()

}

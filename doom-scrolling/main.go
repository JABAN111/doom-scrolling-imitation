package main

import (
	"rshd/lab1/v2/config"
	"rshd/lab1/v2/db"
	"rshd/lab1/v2/logger"
)

func main() {
	log := logger.GetInstance()

	cfg := config.Config{
		CouchBaseCfg: config.CouchBaseConfig{
			URL:      "localhost",
			Username: "jaba_admin",
			Password: "jaba_pwd",
			Bucket:   "doom-scrolling",
		},
	}

	db.InitializeCluster(cfg, log)
	client, err := db.InitClient(cfg, log)
	if err != nil {
		panic(err)
	}
	log.Info("client", "client", client)

}

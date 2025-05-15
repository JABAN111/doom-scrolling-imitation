package document

import (
	"github.com/couchbase/gocb/v2"
	"log/slog"
	"rshd/lab1/v2/config"
)

// InitializeCluster возвращает пропингованный кластер
func InitializeCluster(cfg config.Config, log *slog.Logger) (*gocb.Cluster, error) {
	gocb.SetLogger(gocb.VerboseStdioLogger())
	log.Info("Initializing Couchbase cluster")

	options := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: cfg.CouchBaseCfg.Username,
			Password: cfg.CouchBaseCfg.Password,
		},
	}

	if err := options.ApplyProfile(gocb.ClusterConfigProfileWanDevelopment); err != nil {
		log.Error("Failed to apply Couchbase configuration profile", "error", err)
		return nil, err
	}

	cluster, err := gocb.Connect(cfg.CouchBaseCfg.URL, options)
	if err != nil {
		log.Error("Failed to connect to Couchbase cluster", "error", err)
		return nil, err
	}

	return cluster, nil
}

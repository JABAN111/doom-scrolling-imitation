package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type CouchBaseConfig struct {
	URL      string `yaml:"url" default:"localhost"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Bucket   string `yaml:"bucket" default:"doom-scrolling"`
}

type DgraphConfig struct {
	URL string `yaml:"url" default:"localhost:9080"`
}

type Config struct {
	CouchBaseCfg CouchBaseConfig
	DgraphCfg    DgraphConfig
	// TMP
	Neo4jURI string
}

func MustLoad(configPath string) Config {
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config %q: %s", configPath, err)
	}
	return cfg
}

package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
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
}

func MustLoad(configPath string) Config {
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config %q: %s", configPath, err)
	}
	return cfg
}

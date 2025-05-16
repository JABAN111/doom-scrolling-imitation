package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type CouchBaseConfig struct {
	URL      string `yaml:"url" env:"COUCHBASE_URL" default:"localhost"`
	Username string `yaml:"username" env:"COUCHBASE_USER" default:"jaba_admin"`
	Password string `yaml:"password" env:"COUCHBASE_PWD" default:"jaba_pwd"`
	Bucket   string `yaml:"bucket" env:"COUCHBASE_BUCK" default:"doom-scrolling"`
}

type ClickhouseConfig struct {
	URL string `yaml:"url" env:"CLICKHOUSE_URL" default:"localhost:9080"`
}

type Influx struct {
	URL string `yaml:"url" env:"INFLUX_URL" default:"localhost:8086"`
}

type Neo4j struct {
	URL string `yaml:"url" env:"NEO_URL" default:"localhost:7687"`
}

type Minio struct {
	URL string `yaml:"url" env:"MINIO_URL" default:"localhost:9000"`
}

type Config struct {
	CouchBaseCfg     CouchBaseConfig  `yaml:"couchBaseCfg"`
	ClickhouseConfig ClickhouseConfig `yaml:"clickhouseConfig"`
	Influx           Influx           `yaml:"influx"`
	Neo4j            Neo4j            `yaml:"neo4J"`
	Minio            Minio            `yaml:"minio"`
}

func MustLoad(configPath string) Config {
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config %q: %s", configPath, err)
	}
	return cfg
}

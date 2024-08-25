package config

import (
	"log"
	"os"
	"sync"
)

var (
	conf *Config
	once *sync.Once
)

type MongoDBConfig struct {
	MongoDBInstance string
	MongoDBName     string
}

type Config struct {
	MongoConfig MongoDBConfig
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

func loadConfig() (*Config, error) {
	var config Config

	config.MongoConfig = MongoDBConfig{
		MongoDBInstance: getEnv("MONGO_DB_INSTANCE", "localhost"),
		MongoDBName:     getEnv("MONGO_DB_NAME", "local_db"),
	}

	return &config, nil
}

func GetConfig() *Config {
	once.Do(func() {
		var err error
		conf, err = loadConfig()
		if err != nil {
			log.Fatalf("Failed to load configuration %s", err)
		}
	})

	return conf
}

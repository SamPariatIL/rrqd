package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	AwsConfig     AWSConfig
	FinopsConfig  FinopsConfig
	JwtConfig     JWTConfig
	KaleyraConfig KaleyraConfig
	MongoConfig   MongoDBConfig
}

type AWSConfig struct {
	AccessKey       string
	SecretAccessKey string
	S3BucketName    string
	S3Region        string
}

type FinopsConfig struct {
	LoginUrl                string
	Email                   string
	Password                string
	TransactionStatusUrl    string
	TransactionStatusOffset int
	TransactionStatusLimit  int
	TransactionStatusFinopr string
}

type JWTConfig struct {
	SecretKey    string
	JWTSecretKey string
}

type KaleyraConfig struct {
	SId        string
	APIKey     string
	TemplateId string
	SenderId   string
}

type MongoDBConfig struct {
	MongoDBInstance string
	MongoDBName     string
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		log.Printf("Failed to parse %s, using fallback %d", key, fallback)
	}

	return fallback
}

func loadConfig() (*Config, error) {
	var config Config

	config.AwsConfig = AWSConfig{
		AccessKey:       getEnv("AWS_ACCESS_KEY_ID", ""),
		SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", ""),
		S3BucketName:    getEnv("AWS_S3_BUCKET_NAME", ""),
		S3Region:        getEnv("AWS_S3_REGION", ""),
	}

	config.FinopsConfig = FinopsConfig{
		LoginUrl:                getEnv("FINOPS_LOGIN_URL", ""),
		Email:                   getEnv("FINOPS_EMAIL", ""),
		Password:                getEnv("FINOPS_PASSWORD", ""),
		TransactionStatusUrl:    getEnv("FINOPS_TRANSACTION_STATUS_URL", ""),
		TransactionStatusOffset: getEnvInt("FINOPS_TRANSACTION_STATUS_OFFSET", 0),
		TransactionStatusLimit:  getEnvInt("FINOPS_TRANSACTION_STATUS_LIMIT", 0),
		TransactionStatusFinopr: getEnv("FINOPS_TRANSACTION_STATUS_FINOPR", ""),
	}

	config.JwtConfig = JWTConfig{
		SecretKey:    getEnv("SECRET_KEY", ""),
		JWTSecretKey: getEnv("JWT_SECRET_KEY", ""),
	}

	config.KaleyraConfig = KaleyraConfig{
		APIKey:     getEnv("KALEYRA_API_KEY", ""),
		SenderId:   getEnv("KALEYRA_SENDER_ID", ""),
		SId:        getEnv("KALEYRA_SID", ""),
		TemplateId: getEnv("KALEYRA_TEMPLATE_ID", ""),
	}

	config.MongoConfig = MongoDBConfig{
		MongoDBInstance: getEnv("MONGO_DB_INSTANCE", ""),
		MongoDBName:     getEnv("MONGO_DB_NAME", ""),
	}

	return &config, nil
}

func GetConfig() *Config {
	once.Do(func() {
		var err error

		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Failed to load .env %s", err)
		}

		conf, err = loadConfig()
		if err != nil {
			log.Fatalf("Failed to load configuration %s", err)
		}
	})

	return conf
}

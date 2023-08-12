package config

import (
	"encoding/json"
	"log"
	"os"

	configdev "Golang/config/dev"
	configLocal "Golang/config/local"
	configprod "Golang/config/production"
	configStg "Golang/config/stag"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env        string           `json:"env"`
	Port       string           `json:"port"`
	Postgresql PostgresqlConfig `json:"postgresql"`
	Redis      RedisConfig      `json:"redis"`
	Aws        AwsConfig
}
type PostgresqlConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"password"`
	Db   string `json:"db"`
}
type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Pass     string `json:"password"`
	Username string `json:"Username"`
	Db       string `json:"db"`
}
type AwsConfig struct {
	Host   string
	Port   string
	Region string
	Id     string
	Secret string
	Token  string
}

func Getconfig() AppConfig {
	var appConfig AppConfig
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	env := os.Getenv("ENVIRONMENT")
	if env == "local" {
		err := json.Unmarshal([]byte(configLocal.ConfigApp), &appConfig)
		if err != nil {
			panic(err)
		}
	} else if env == "dev" {
		err := json.Unmarshal([]byte(configdev.ConfigApp), &appConfig)
		if err != nil {
			panic(err)
		}
	} else if env == "stg" {
		err := json.Unmarshal([]byte(configStg.ConfigApp), &appConfig)
		if err != nil {
			panic(err)
		}
	} else if env == "prod" {
		err := json.Unmarshal([]byte(configprod.ConfigApp), &appConfig)
		if err != nil {
			panic(err)
		}
	}
	appConfig.Aws.Host = os.Getenv("AWS_HOST")
	appConfig.Aws.Port = os.Getenv("AWS_PORT")
	appConfig.Aws.Region = os.Getenv("AWS_REGION")
	appConfig.Aws.Id = os.Getenv("AWS_ID")
	appConfig.Aws.Secret = os.Getenv("AWS_SECRET")
	appConfig.Aws.Token = os.Getenv("AWS_TOKEN")
	if appConfig.Port == "" {
		appConfig.Port = os.Getenv("PORT")
	}
	//MYSQL
	if appConfig.Postgresql.User == "" {
		appConfig.Postgresql.User = os.Getenv("DB_USER")

	}
	if appConfig.Postgresql.Pass == "" {
		appConfig.Postgresql.Pass = os.Getenv("DB_PASSWORD")

	}
	if appConfig.Postgresql.Host == "" {
		appConfig.Postgresql.Host = os.Getenv("DB_HOST")

	}
	if appConfig.Postgresql.Port == "" {
		appConfig.Postgresql.Port = os.Getenv("DB_PORT")

	}
	if appConfig.Postgresql.Db == "" {
		appConfig.Postgresql.Db = os.Getenv("DB_NAME")

	}
	//REDIS
	if appConfig.Redis.Host == "" {
		appConfig.Redis.Host = os.Getenv("REDIS_HOST")

	}
	if appConfig.Redis.Port == "" {
		appConfig.Redis.Port = os.Getenv("REDIS_PORT")

	}
	if appConfig.Redis.Pass == "" {
		appConfig.Redis.Pass = os.Getenv("REDIS_PASSWORD")

	}
	if appConfig.Redis.Username == "" {
		appConfig.Redis.Username = os.Getenv("REDIS_USER")

	}
	if appConfig.Redis.Db == "" {
		appConfig.Redis.Db = os.Getenv("REDIS_DB")

	}
	return appConfig
}

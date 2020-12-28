package config

import (
	"log"
	"os"
)

type DB struct {
	Driver    string
	User      string
	Password  string
	Name      string
	Address   string
	SourceURL string
}
type Server struct {
	Env          string
	Port         string
	JwtSecretKey string
}

type Config struct {
	DB
	Server
}

var Cfg Config

func init() {
	Cfg = Config{
		DB: DB{
			Driver:    GetEnv("DB_DRIVER"),
			User:      GetEnv("DB_USER"),
			Password:  GetEnv("DB_PASSWORD"),
			Name:      GetEnv("DB_NAME"),
			Address:   GetEnv("DB_ADDRESS"),
			SourceURL: GetEnv("MIGRATE_SOURCE_FILE"),
		},
		Server: Server{
			Env:          GetEnv("ENV"),
			Port:         GetEnv("PORT"),
			JwtSecretKey: GetEnv("JWT_SECRET_KEY"),
		},
	}
}

func GetEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		log.Fatalf("action=get env variable, err=%s is not set.", key)
	}
	return env
}

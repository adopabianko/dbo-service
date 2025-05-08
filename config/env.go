package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	EnvLocal      = "local"
	EnvProduction = "production"
	EnvStaging    = "development"
	EnvUAT        = "uat"
)

type Provider struct {
	App        *App
	JWT        *JWT
	Postgresql *Postgresql
}

type App struct {
	Env  string
	Host string
	Port int
	Name string
}

type Postgresql struct {
	DBConnection      string
	DBMaxConns        int
	DBMaxConnLifetime int
	DBMaxConnIdletime int
}

type JWT struct {
	JWTSecret string
}

func InitConfig() (*Provider, error) {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}

	conf := &Provider{
		App: &App{
			Env:  getEnv("APP_ENV", ""),
			Host: getEnv("APP_HOST", ""),
			Port: getEnvAsInt("APP_PORT", 8080),
			Name: getEnv("APP_NAME", ""),
		},
		Postgresql: &Postgresql{
			DBConnection:      getEnv("DB_CONNECTION", ""),
			DBMaxConns:        getEnvAsInt("DB_MAX_CONNS", 20),
			DBMaxConnLifetime: getEnvAsInt("DB_MAX_CONN_LIFETIME", 30),
			DBMaxConnIdletime: getEnvAsInt("DB_MAX_CONN_IDLETIME", 10),
		},
		JWT: &JWT{
			JWTSecret: getEnv("JWT_SECRET", ""),
		},
	}

	return conf, nil
}

func getEnvAsInt(name string, defaultVal int) int {
	valStr := getEnv(name, "")
	if value, err := strconv.Atoi(valStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

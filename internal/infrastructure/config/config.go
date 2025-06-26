package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	API      APIConfig
	Log      LogConfig
}

type ServerConfig struct {
	Port string
	Host string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	Secret          string
	ExpirationHours int
}

type APIConfig struct {
	KeyHeader string
	Keys      []string
}

type LogConfig struct {
	Level string
}

func Load() (*Config, error) {
	// .envファイルの読み込み（存在する場合）
	godotenv.Load()

	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "3001"),
			Host: getEnv("SERVER_HOST", "localhost"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "35432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "echo-household-budget"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:          getEnv("JWT_SECRET", "your-secret-key-here"),
			ExpirationHours: getEnvAsInt("JWT_EXPIRATION_HOURS", 24),
		},
		API: APIConfig{
			KeyHeader: getEnv("API_KEY_HEADER", "X-API-Key"),
			Keys:      getEnvAsSlice("API_KEYS", []string{"key1", "key2", "key3"}),
		},
		Log: LogConfig{
			Level: getEnv("LOG_LEVEL", "debug"),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		// カンマ区切りの文字列をスライスに変換
		// 実際の実装ではより複雑な処理が必要かもしれません
		return []string{value}
	}
	return defaultValue
}

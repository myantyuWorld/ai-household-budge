package main

import (
	"log"
	"os"

	"ai-household-budge/internal/infrastructure/config"
	"ai-household-budge/internal/infrastructure/server"
)

func main() {
	// 設定の読み込み
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// サーバーの初期化
	srv := server.NewServer(cfg)

	// サーバーの起動
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3001"
	}

	log.Printf("Server starting on port %s", port)
	if err := srv.Start(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

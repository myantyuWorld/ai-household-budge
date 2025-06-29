package server

import (
	"ai-household-budge/internal/infrastructure/config"
	"ai-household-budge/internal/infrastructure/middleware"
	"ai-household-budge/internal/presentation/handler"
	"fmt"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	echo *echo.Echo
	cfg  *config.Config
	db   *gorm.DB
}

func NewServer(cfg *config.Config) *Server {
	e := echo.New()

	// ミドルウェアの設定
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS())

	// データベース接続の初期化
	db, err := initDatabase(cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// APIキー認証ミドルウェア
	apiKeyMiddleware := middleware.NewAPIKeyMiddleware(cfg.API)

	server := &Server{
		echo: e,
		cfg:  cfg,
		db:   db,
	}

	// ルーティングの設定
	server.setupRoutes(apiKeyMiddleware)

	return server
}

// initDatabase データベース接続を初期化
func initDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// データベース接続をテスト
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Server) setupRoutes(apiKeyMiddleware *middleware.APIKeyMiddleware) {
	// ヘルスチェックエンドポイント（認証不要）
	s.echo.GET("/health", handler.HealthCheck)

	// APIグループ（認証必要）
	api := s.echo.Group("/api/v1")
	api.Use(apiKeyMiddleware.Authenticate)

	// チャット分析関連のエンドポイント
	// chatHandler := handler.NewChatHandler(s.createChatUseCase())
	// api.POST("/chat/analyze", chatHandler.HandleChat)
	// api.GET("/chat/health", chatHandler.HandleHealth)
}

// func (s *Server) createChatUseCase() usecase.ChatUseCase {
// 	// 分析リポジトリの初期化
// 	// analysisRepo := persistence.NewAnalysisRepository(s.db)
// 	// return usecase.NewChatUseCase(analysisRepo)
// 	return nil
// }

func (s *Server) Start(address string) error {
	return s.echo.Start(address)
}

func (s *Server) Shutdown() error {
	if s.db != nil {
		sqlDB, err := s.db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
	return s.echo.Shutdown(nil)
}

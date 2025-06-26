package server

import (
	"ai-household-budge/internal/infrastructure/config"
	"ai-household-budge/internal/infrastructure/middleware"
	"ai-household-budge/internal/infrastructure/persistence"
	"ai-household-budge/internal/presentation/handler"
	"ai-household-budge/internal/usecase"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
	cfg  *config.Config
}

func NewServer(cfg *config.Config) *Server {
	e := echo.New()

	// ミドルウェアの設定
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS())

	// APIキー認証ミドルウェア
	apiKeyMiddleware := middleware.NewAPIKeyMiddleware(cfg.API)

	server := &Server{
		echo: e,
		cfg:  cfg,
	}

	// ルーティングの設定
	server.setupRoutes(apiKeyMiddleware)

	return server
}

func (s *Server) setupRoutes(apiKeyMiddleware *middleware.APIKeyMiddleware) {
	// ヘルスチェックエンドポイント（認証不要）
	s.echo.GET("/health", handler.HealthCheck)

	// APIグループ（認証必要）
	api := s.echo.Group("/api/v1")
	api.Use(apiKeyMiddleware.Authenticate)

	// カテゴリ関連のエンドポイント
	categoryHandler := handler.NewCategoryHandler(s.createCategoryUseCase())
	api.GET("/categories", categoryHandler.GetAll)
	api.GET("/categories/:id", categoryHandler.GetByID)
	api.POST("/categories", categoryHandler.Create)
	api.PUT("/categories/:id", categoryHandler.Update)
	api.DELETE("/categories/:id", categoryHandler.Delete)
}

func (s *Server) createCategoryUseCase() *usecase.CategoryUseCase {
	// リポジトリの初期化
	repo := persistence.NewCategoryRepository()
	return usecase.NewCategoryUseCase(repo)
}

func (s *Server) Start(address string) error {
	return s.echo.Start(address)
}

func (s *Server) Shutdown() error {
	return s.echo.Shutdown(nil)
}

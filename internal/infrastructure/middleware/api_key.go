package middleware

import (
	"ai-household-budge/internal/infrastructure/config"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type APIKeyMiddleware struct {
	config config.APIConfig
}

func NewAPIKeyMiddleware(cfg config.APIConfig) *APIKeyMiddleware {
	return &APIKeyMiddleware{
		config: cfg,
	}
}

func (m *APIKeyMiddleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// APIキーの取得
		apiKey := c.Request().Header.Get(m.config.KeyHeader)
		if apiKey == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "API key is required")
		}

		// APIキーの検証
		if !m.isValidAPIKey(apiKey) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid API key")
		}

		return next(c)
	}
}

func (m *APIKeyMiddleware) isValidAPIKey(apiKey string) bool {
	// 前後の空白を削除
	apiKey = strings.TrimSpace(apiKey)

	// 設定されたAPIキーと比較
	for _, validKey := range m.config.Keys {
		if apiKey == validKey {
			return true
		}
	}

	return false
}

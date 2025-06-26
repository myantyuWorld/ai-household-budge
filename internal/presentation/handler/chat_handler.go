package handler

import (
	"net/http"

	"ai-household-budge/internal/usecase"

	"github.com/labstack/echo/v4"
)

type (
	ChatRequest struct {
		Message string `json:"message" validate:"required"`
	}

	ChatResponse struct {
		Message   string `json:"message"`
		Analysis  string `json:"analysis"`
		SQLQuery  string `json:"sql_query,omitempty"`
		Timestamp string `json:"timestamp"`
	}

	chatHandler struct {
		chatUseCase usecase.ChatUseCase
	}

	ChatHandler interface {
		HandleChat(c echo.Context) error
		HandleHealth(c echo.Context) error
	}
)

func NewChatHandler(chatUseCase usecase.ChatUseCase) ChatHandler {
	return &chatHandler{
		chatUseCase: chatUseCase,
	}
}

// HandleChat チャットメッセージを処理し、データベース分析結果を返す
func (h *chatHandler) HandleChat(c echo.Context) error {
	var req ChatRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Validation failed")
	}

	// ユースケースを呼び出して分析を実行
	result, err := h.chatUseCase.AnalyzeDatabase(req.Message)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Analysis failed: "+err.Error())
	}

	response := ChatResponse{
		Message:   result.Message,
		Analysis:  result.Analysis,
		SQLQuery:  result.SQLQuery,
		Timestamp: result.Timestamp,
	}

	return c.JSON(http.StatusOK, response)
}

// HandleHealth チャットサービスのヘルスチェック
func (h *chatHandler) HandleHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "healthy",
		"service": "chat-analysis",
	})
}

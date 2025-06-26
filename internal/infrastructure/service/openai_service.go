package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// OpenAIService OpenAI APIを使用したサービスの構造体
type OpenAIService struct {
	apiKey string
	model  string
	client *http.Client
}

// OpenAIRequest OpenAI APIリクエストの構造体
type OpenAIRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// OpenAIResponse OpenAI APIレスポンスの構造体
type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
	Error   *Error   `json:"error,omitempty"`
}

// Message メッセージの構造体
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Choice 選択肢の構造体
type Choice struct {
	Message Message `json:"message"`
}

// Error エラーの構造体
type Error struct {
	Message string `json:"message"`
}

// NewOpenAIService OpenAIサービスの新しいインスタンスを作成
func NewOpenAIService() *OpenAIService {
	return &OpenAIService{
		apiKey: os.Getenv("OPENAI_API_KEY"),
		model:  getEnv("OPENAI_MODEL", "gpt-3.5-turbo"),
		client: &http.Client{},
	}
}

// ConvertToSQL 自然言語をSQLに変換
func (s *OpenAIService) ConvertToSQL(message string, schema string) (string, error) {
	if s.apiKey == "" {
		return "", fmt.Errorf("OpenAI API key is not set")
	}

	prompt := fmt.Sprintf(`あなたはSQLデータベースの専門家です。以下のデータベーススキーマに基づいて、自然言語の質問をSQLクエリに変換してください。

データベーススキーマ:
%s

質問: %s

SQLクエリのみを返してください。説明は不要です。`, schema, message)

	requestBody := OpenAIRequest{
		Model: s.model,
		Messages: []Message{
			{
				Role:    "system",
				Content: "あなたはSQLデータベースの専門家です。自然言語の質問をSQLクエリに変換してください。",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.apiKey)

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if response.Error != nil {
		return "", fmt.Errorf("OpenAI API error: %s", response.Error.Message)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI API")
	}

	return response.Choices[0].Message.Content, nil
}

// AnalyzeResults データベース結果を自然言語で分析
func (s *OpenAIService) AnalyzeResults(originalMessage string, data interface{}) (string, error) {
	if s.apiKey == "" {
		return "", fmt.Errorf("OpenAI API key is not set")
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	prompt := fmt.Sprintf(`以下のデータベースクエリ結果を分析し、自然言語で説明してください。

元の質問: %s
データベース結果: %s

分析結果を日本語で分かりやすく説明してください。`, originalMessage, string(dataJSON))

	requestBody := OpenAIRequest{
		Model: s.model,
		Messages: []Message{
			{
				Role:    "system",
				Content: "あなたはデータ分析の専門家です。データベースの結果を自然言語で分かりやすく説明してください。",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.apiKey)

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if response.Error != nil {
		return "", fmt.Errorf("OpenAI API error: %s", response.Error.Message)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI API")
	}

	return response.Choices[0].Message.Content, nil
}

// getEnv 環境変数を取得（デフォルト値付き）
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

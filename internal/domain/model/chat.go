package model

type Chat struct {
	message string
}

func NewChat(message string) *Chat {
	return &Chat{
		message: message,
	}
}

func (c *Chat) GetMessage() string {
	return c.message
}

func (c *Chat) FallbackMessage() string {
	return "データベースの分析結果をお返しします。詳細な分析については、より高度なAI機能の実装が必要です。"
}

func (c *Chat) FallbackConvertToSQL() string {
	// 簡単なマッピングで実装
	switch c.message {
	case "分析履歴を教えて":
		return "SELECT * FROM analysis_history ORDER BY created_at DESC LIMIT 10"
	case "最近の分析を教えて":
		return "SELECT message, analysis, timestamp FROM analysis_history ORDER BY timestamp DESC LIMIT 5"
	default:
		return "SELECT * FROM analysis_history LIMIT 10"
	}
}

func (c *Chat) GetDatabaseSchema() string {
	// 実際のデータベーススキーマを返す
	// 現在は仮のスキーマ
	return `
		-- 分析履歴テーブル
		CREATE TABLE analysis_history (
			id VARCHAR(36) PRIMARY KEY,
			user_id VARCHAR(255) NOT NULL,
			message TEXT NOT NULL,
			sql_query TEXT NOT NULL,
			analysis TEXT NOT NULL,
			timestamp TIMESTAMP NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
}

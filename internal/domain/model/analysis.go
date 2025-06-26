package model

import (
	"time"

	"gorm.io/gorm"
)

// AnalysisHistory 分析履歴のドメインモデル
type AnalysisHistory struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	UserID    string         `json:"user_id" gorm:"type:varchar(255);not null"`
	Message   string         `json:"message" gorm:"type:text;not null"`
	SQLQuery  string         `json:"sql_query" gorm:"type:text;not null"`
	Analysis  string         `json:"analysis" gorm:"type:text;not null"`
	Timestamp time.Time      `json:"timestamp" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName テーブル名を指定
func (AnalysisHistory) TableName() string {
	return "analysis_history"
}

// AnalysisResult 分析結果のドメインモデル
type AnalysisResult struct {
	Data      interface{} `json:"data"`
	RowCount  int         `json:"row_count"`
	Error     string      `json:"error,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

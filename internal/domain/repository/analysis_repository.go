package repository

import (
	"ai-household-budge/internal/domain/model"
)

// AnalysisRepository 分析リポジトリのインターフェース
type AnalysisRepository interface {
	// ExecuteQuery SQLクエリを実行してデータを取得
	ExecuteQuery(sqlQuery string) (*model.AnalysisResult, error)

	// SaveAnalysis 分析履歴を保存
	SaveAnalysis(analysis *model.AnalysisHistory) error

	// GetAnalysisHistory ユーザーの分析履歴を取得
	GetAnalysisHistory(userID string, limit int) ([]*model.AnalysisHistory, error)

	// GetAnalysisByID 特定の分析履歴を取得
	GetAnalysisByID(id string) (*model.AnalysisHistory, error)
}

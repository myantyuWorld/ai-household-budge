package persistence

import (
	"ai-household-budge/internal/domain/model"
	"ai-household-budge/internal/domain/repository"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// analysisRepository 分析リポジトリの実装
type analysisRepository struct {
	db *gorm.DB
}

// NewAnalysisRepository 分析リポジトリの新しいインスタンスを作成
func NewAnalysisRepository(db *gorm.DB) repository.AnalysisRepository {
	return &analysisRepository{
		db: db,
	}
}

// ExecuteQuery SQLクエリを実行してデータを取得
func (r *analysisRepository) ExecuteQuery(sqlQuery string) (*model.AnalysisResult, error) {
	var results []map[string]interface{}

	// GORMのRawメソッドを使用してSQLクエリを実行
	err := r.db.Raw(sqlQuery).Scan(&results).Error
	if err != nil {
		return &model.AnalysisResult{
			Error:     err.Error(),
			Timestamp: time.Now(),
		}, nil
	}

	return &model.AnalysisResult{
		Data:      results,
		RowCount:  len(results),
		Timestamp: time.Now(),
	}, nil
}

// SaveAnalysis 分析履歴を保存
func (r *analysisRepository) SaveAnalysis(analysis *model.AnalysisHistory) error {
	analysis.ID = uuid.New().String()
	analysis.Timestamp = time.Now()

	// GORMのCreateメソッドを使用してレコードを作成
	return r.db.Create(analysis).Error
}

// GetAnalysisHistory ユーザーの分析履歴を取得
func (r *analysisRepository) GetAnalysisHistory(userID string, limit int) ([]*model.AnalysisHistory, error) {
	var histories []*model.AnalysisHistory

	// GORMのWhereとLimitメソッドを使用してクエリを実行
	err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Find(&histories).Error

	return histories, err
}

// GetAnalysisByID 特定の分析履歴を取得
func (r *analysisRepository) GetAnalysisByID(id string) (*model.AnalysisHistory, error) {
	var history model.AnalysisHistory

	// GORMのFirstメソッドを使用してレコードを取得
	err := r.db.Where("id = ?", id).First(&history).Error
	if err != nil {
		return nil, err
	}

	return &history, nil
}

package usecase

import (
	"ai-household-budge/internal/domain/model"
	"ai-household-budge/internal/domain/repository"
	"ai-household-budge/internal/infrastructure/service"
	"time"
)

type (
	ChatAnalysisInput struct {
		Message string
	}

	ChatAnalysisOutput struct {
		Message   string
		Analysis  string
		SQLQuery  string
		Timestamp string
	}

	chatUseCase struct {
		analysisRepository repository.AnalysisRepository
		openAIService      *service.OpenAIService
	}

	// ChatUseCase チャットユースケースのインターフェース
	ChatUseCase interface {
		AnalyzeDatabase(message string) (*ChatAnalysisOutput, error)
	}
)

// NewChatUseCase チャットユースケースの新しいインスタンスを作成
func NewChatUseCase(analysisRepository repository.AnalysisRepository) ChatUseCase {
	return &chatUseCase{
		analysisRepository: analysisRepository,
		openAIService:      service.NewOpenAIService(),
	}
}

// AnalyzeDatabase 自然言語メッセージを分析し、データベース分析結果を返す
func (uc *chatUseCase) AnalyzeDatabase(message string) (*ChatAnalysisOutput, error) {
	chat := model.NewChat(message)

	// 1. データベーススキーマを取得
	schema := chat.GetDatabaseSchema()

	// 2. 自然言語をSQLに変換
	sqlQuery, err := uc.convertToSQL(chat, schema)
	if err != nil {
		return nil, err
	}

	// 3. SQLを実行してデータを取得
	result, err := uc.analysisRepository.ExecuteQuery(sqlQuery)
	if err != nil {
		return nil, err
	}

	// 4. 結果を自然言語で分析
	analysis, err := uc.analyzeResults(chat, result.Data)
	if err != nil {
		return nil, err
	}

	// 5. 分析履歴を保存
	err = uc.analysisRepository.SaveAnalysis(&model.AnalysisHistory{
		Message:   message,
		SQLQuery:  sqlQuery,
		Analysis:  analysis,
		Timestamp: time.Now(),
	})
	if err != nil {
		// ログに記録するが、エラーは返さない
		// TODO: ログ実装
	}

	output := uc.makeOutput(chat, analysis, sqlQuery)

	return output, nil
}

func (uc *chatUseCase) makeOutput(chat *model.Chat, analysis string, sqlQuery string) *ChatAnalysisOutput {
	return &ChatAnalysisOutput{
		Message:   chat.GetMessage(),
		Analysis:  analysis,
		SQLQuery:  sqlQuery,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

// convertToSQL 自然言語をSQLに変換
func (uc *chatUseCase) convertToSQL(chat *model.Chat, schema string) (string, error) {
	// OpenAI APIを使用して自然言語をSQLに変換
	sqlQuery, err := uc.openAIService.ConvertToSQL(chat.GetMessage(), schema)
	if err != nil {
		// OpenAI APIが利用できない場合は、フォールバックとして簡単なマッピングを使用
		return chat.FallbackConvertToSQL(), nil
	}

	return sqlQuery, nil
}

// analyzeResults データベース結果を自然言語で分析
func (uc *chatUseCase) analyzeResults(chat *model.Chat, data interface{}) (string, error) {
	// OpenAI APIを使用してデータを自然言語で分析
	analysis, err := uc.openAIService.AnalyzeResults(chat.GetMessage(), data)
	if err != nil {
		// OpenAI APIが利用できない場合は、フォールバックとして簡単なテンプレートを使用
		return chat.FallbackMessage(), nil
	}

	return analysis, nil
}

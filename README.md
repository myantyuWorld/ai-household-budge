# AI Household Budget - 自然言語によるデータベース分析 API

買い物メモアプリから利用する自然言語によるデータベース分析 API です。Go/Echo を使用したレイヤードアーキテクチャで構築されており、OpenAI API を活用して自然言語の質問を SQL クエリに変換し、データベース分析結果を自然言語で返します。

## アーキテクチャ

このプロジェクトは以下のレイヤードアーキテクチャに従っています：

```
┌─────────────────────────────────────┐
│           Presentation Layer        │
│         (Handlers, Routes)          │
├─────────────────────────────────────┤
│            Use Case Layer           │
│        (Business Logic)             │
├─────────────────────────────────────┤
│           Domain Layer              │
│    (Models, Repository Interfaces)  │
├─────────────────────────────────────┤
│        Infrastructure Layer         │
│   (Repository Implementations,      │
│    Database, External Services)     │
└─────────────────────────────────────┘
```

## 機能

- **自然言語によるデータベース分析**: 自然言語の質問を SQL クエリに変換し、分析結果を自然言語で返す
- **OpenAI API 統合**: GPT-3.5-turbo を使用した高度な自然言語処理
- **分析履歴管理**: ユーザーの分析履歴を保存・管理
- **API キー認証**: セキュアな API アクセス制御
- **ヘルスチェック**: サービス状態の監視

## 技術スタック

- **言語**: Go 1.21
- **Web フレームワーク**: Echo v4
- **認証**: API キー認証
- **データベース**: PostgreSQL
- **AI サービス**: OpenAI API (GPT-3.5-turbo)
- **コンテナ**: Docker
- **オーケストレーション**: Docker Compose

## セットアップ

### 前提条件

- Go 1.21 以上
- Docker & Docker Compose
- OpenAI API キー

### ローカル開発

1. リポジトリをクローン

```bash
git clone <repository-url>
cd ai-household-budge
```

2. 依存関係をインストール

```bash
go mod download
```

3. 環境変数を設定

```bash
cp env.example .env
# .envファイルを編集して必要な設定を行ってください
# 特にOPENAI_API_KEYの設定が重要です
```

4. アプリケーションを起動

```bash
go run cmd/web_api/main.go
```

### Docker を使用した開発

1. Docker Compose で起動

```bash
docker-compose up --build
```

2. アプリケーションにアクセス

```
http://localhost:3001
```

## API エンドポイント

### 認証

すべての API エンドポイント（`/health`を除く）は API キー認証が必要です。

**ヘッダー**: `X-API-Key: your-api-key`

### エンドポイント一覧

#### ヘルスチェック

- `GET /health` - サービス状態の確認

#### チャット分析

- `POST /api/v1/chat/analyze` - 自然言語によるデータベース分析
- `GET /api/v1/chat/health` - チャットサービスのヘルスチェック

## 使用例

### 自然言語によるデータベース分析

```bash
curl -X POST http://localhost:3001/api/v1/chat/analyze \
  -H "X-API-Key: key1" \
  -H "Content-Type: application/json" \
  -d '{
    "message": "分析履歴を教えて",
    "user_id": "user123"
  }'
```

**レスポンス例:**

```json
{
  "message": "分析履歴を教えて",
  "analysis": "分析履歴の結果をお伝えします。最近の分析では、データベースの構造に関する質問が多く見られます。",
  "sql_query": "SELECT * FROM analysis_history ORDER BY created_at DESC LIMIT 10",
  "timestamp": "2024-01-15T10:30:00Z"
}
```

## 環境変数

| 変数名           | 説明                            | デフォルト値         |
| ---------------- | ------------------------------- | -------------------- |
| `SERVER_PORT`    | サーバーポート                  | 3001                 |
| `SERVER_HOST`    | サーバーホスト                  | localhost            |
| `DB_HOST`        | データベースホスト              | localhost            |
| `DB_PORT`        | データベースポート              | 5432                 |
| `DB_USER`        | データベースユーザー            | postgres             |
| `DB_PASSWORD`    | データベースパスワード          | password             |
| `DB_NAME`        | データベース名                  | ai_household_budge   |
| `JWT_SECRET`     | JWT シークレットキー            | your-secret-key-here |
| `API_KEY_HEADER` | API キーヘッダー名              | X-API-Key            |
| `API_KEYS`       | 有効な API キー（カンマ区切り） | key1,key2,key3       |
| `OPENAI_API_KEY` | OpenAI API キー                 | 必須                 |
| `OPENAI_MODEL`   | OpenAI モデル名                 | gpt-3.5-turbo        |

## 開発

### プロジェクト構造

```
ai-household-budge/
├── cmd/
│   └── web_api/
│       └── main.go                 # エントリーポイント
├── internal/
│   ├── domain/                     # ドメイン層
│   │   ├── model/                  # ドメインモデル
│   │   │   └── analysis.go
│   │   └── repository/             # リポジトリインターフェース
│   │       └── analysis_repository.go
│   ├── infrastructure/             # インフラストラクチャ層
│   │   ├── config/                 # 設定管理
│   │   ├── middleware/             # ミドルウェア
│   │   ├── persistence/            # データ永続化
│   │   │   └── analysis_repository.go
│   │   ├── service/                # 外部サービス
│   │   │   └── openai_service.go
│   │   └── server/                 # サーバー設定
│   ├── presentation/               # プレゼンテーション層
│   │   └── handler/                # HTTPハンドラー
│   │       ├── chat_handler.go
│   │       └── health.go
│   └── usecase/                    # ユースケース層
│       └── chat_usecase.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

### テスト

```bash
go test ./...
```

## 自然言語分析の仕組み

この API は以下の流れで自然言語によるデータベース分析を実行します：

1. **自然言語入力**: ユーザーから自然言語の質問を受け取る
2. **SQL 変換**: OpenAI API を使用して自然言語を SQL クエリに変換
3. **データ取得**: 生成された SQL クエリをデータベースで実行
4. **結果分析**: 取得したデータを自然言語で分析・説明
5. **履歴保存**: 分析履歴をデータベースに保存

### 対応している質問例

- "分析履歴を教えて"
- "最近の分析を教えて"
- "特定のユーザーの分析履歴を教えて"
- "分析結果の統計を教えて"

## ライセンス

このプロジェクトは MIT ライセンスの下で公開されています。

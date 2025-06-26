# AI Household Budget - カテゴリ管理マイクロサービス

買い物メモアプリから利用するカテゴリ管理マイクロサービスです。Go/Echo を使用したレイヤードアーキテクチャで構築されています。

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

- **カテゴリ管理**: カテゴリの作成、取得、更新、削除
- **API キー認証**: セキュアな API アクセス制御
- **ヘルスチェック**: サービス状態の監視

## 技術スタック

- **言語**: Go 1.21
- **Web フレームワーク**: Echo v4
- **認証**: API キー認証
- **データベース**: PostgreSQL (将来的に実装予定)
- **コンテナ**: Docker
- **オーケストレーション**: Docker Compose

## セットアップ

### 前提条件

- Go 1.21 以上
- Docker & Docker Compose

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
```

4. アプリケーションを起動

```bash
go run cmd/main.go
```

### Docker を使用した開発

1. Docker Compose で起動

```bash
docker-compose up --build
```

2. アプリケーションにアクセス

```
http://localhost:8080
```

## API エンドポイント

### 認証

すべての API エンドポイント（`/health`を除く）は API キー認証が必要です。

**ヘッダー**: `X-API-Key: your-api-key`

### エンドポイント一覧

#### ヘルスチェック

- `GET /health` - サービス状態の確認

#### カテゴリ

- `GET /api/v1/categories` - 全カテゴリの取得
- `GET /api/v1/categories/:id` - 特定のカテゴリの取得
- `POST /api/v1/categories` - カテゴリの作成
- `PUT /api/v1/categories/:id` - カテゴリの更新
- `DELETE /api/v1/categories/:id` - カテゴリの削除

## 使用例

### カテゴリの作成

```bash
curl -X POST http://localhost:8080/api/v1/categories \
  -H "X-API-Key: key1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "食料品",
    "description": "食品や飲料",
    "color": "#FF6B6B"
  }'
```

### カテゴリの取得

```bash
curl -X GET http://localhost:8080/api/v1/categories \
  -H "X-API-Key: key1"
```

### カテゴリの更新

```bash
curl -X PUT http://localhost:8080/api/v1/categories/category-id \
  -H "X-API-Key: key1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "食品",
    "description": "食品や飲料のカテゴリ",
    "color": "#4ECDC4"
  }'
```

## 環境変数

| 変数名           | 説明                            | デフォルト値         |
| ---------------- | ------------------------------- | -------------------- |
| `SERVER_PORT`    | サーバーポート                  | 8080                 |
| `SERVER_HOST`    | サーバーホスト                  | localhost            |
| `DB_HOST`        | データベースホスト              | localhost            |
| `DB_PORT`        | データベースポート              | 5432                 |
| `DB_USER`        | データベースユーザー            | postgres             |
| `DB_PASSWORD`    | データベースパスワード          | password             |
| `DB_NAME`        | データベース名                  | ai_household_budge   |
| `JWT_SECRET`     | JWT シークレットキー            | your-secret-key-here |
| `API_KEY_HEADER` | API キーヘッダー名              | X-API-Key            |
| `API_KEYS`       | 有効な API キー（カンマ区切り） | key1,key2,key3       |

## 開発

### プロジェクト構造

```
ai-household-budge/
├── cmd/
│   └── main.go                 # エントリーポイント
├── internal/
│   ├── domain/                 # ドメイン層
│   │   ├── model/             # ドメインモデル
│   │   └── repository/        # リポジトリインターフェース
│   ├── infrastructure/        # インフラストラクチャ層
│   │   ├── config/           # 設定管理
│   │   ├── middleware/       # ミドルウェア
│   │   ├── persistence/      # データ永続化
│   │   └── server/           # サーバー設定
│   ├── presentation/         # プレゼンテーション層
│   │   └── handler/          # HTTPハンドラー
│   └── usecase/             # ユースケース層
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

## ライセンス

このプロジェクトは MIT ライセンスの下で公開されています。

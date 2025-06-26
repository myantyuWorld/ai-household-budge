#!/bin/bash

# APIテストスクリプト
BASE_URL="http://localhost:3001"
API_KEY="key1"

echo "=== AI Household Budget API テスト ==="

# ヘルスチェック
echo "1. ヘルスチェック"
curl -s -X GET "$BASE_URL/health" | jq .

echo -e "\n2. チャットヘルスチェック"
curl -s -X GET "$BASE_URL/api/v1/chat/health" \
  -H "X-API-Key: $API_KEY" | jq .

echo -e "\n3. 自然言語によるデータベース分析テスト"
curl -s -X POST "$BASE_URL/api/v1/chat/analyze" \
  -H "X-API-Key: $API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "message": "分析履歴を教えて",
    "user_id": "test_user_123"
  }' | jq .

echo -e "\n=== テスト完了 ===" 

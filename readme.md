# メモ帳 API

シンプルなメモ帳 API。タイトル、本文、タグでメモを管理できます。

## 機能

- メモの作成、取得、更新、削除（CRUD 操作）
- ローカル JSON ファイルによるデータ永続化

## データ構造

```json
{
  "memos": [
    {
      "id": "uuid",
      "title": "メモのタイトル",
      "content": "メモの本文",
      "created_at": "2024-03-21T10:00:00Z",
      "updated_at": "2024-03-21T10:00:00Z"
    }
  ]
}
```

## API エンドポイント

### メモ関連

- `GET /memos` - 全メモの取得
- `GET /memos/:id` - 特定のメモの取得
- `POST /memos` - 新規メモの作成
- `PUT /memos/:id` - メモの更新
- `DELETE /memos/:id` - メモの削除

## 技術スタック

- Go 1.22
- 標準ライブラリのみを使用
- JSON ファイルによるデータ永続化

## 開発環境のセットアップ

1. Go のインストール
2. リポジトリのクローン
   ```bash
   git clone https://github.com/masvc/go_memo.git
   cd go_memo
   ```
3. 依存関係のインストール
   ```bash
   go mod tidy
   ```
4. サーバーの起動
   ```bash
   go run main.go
   ```

サーバーは `http://localhost:8080` で起動します。

## 今後の拡張案

- タグ機能の実装
- メモの検索機能
- メモの並び替え機能
- メモのカテゴリ分け
- メモのエクスポート/インポート機能
- メモの共有機能

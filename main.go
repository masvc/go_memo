package main

import (
	"log"
	"memo-api/handlers"
	"net/http"
)

func main() {
	// メモ関連のエンドポイントを設定
	// GET /memos - 全メモの取得
	// POST /memos - 新規メモの作成
	http.HandleFunc("/memos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetAllMemos(w, r)
		case http.MethodPost:
			handlers.CreateMemo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// 特定のメモに対するエンドポイントを設定
	// GET /memos/:id - 特定のメモの取得
	// PUT /memos/:id - メモの更新
	// DELETE /memos/:id - メモの削除
	http.HandleFunc("/memos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetMemoByID(w, r)
		case http.MethodPut:
			handlers.UpdateMemo(w, r)
		case http.MethodDelete:
			handlers.DeleteMemo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// サーバーを起動
	log.Println("メモ帳APIサーバーを起動します...")
	log.Println("エンドポイント:")
	log.Println("  GET    /memos     - 全メモの取得")
	log.Println("  POST   /memos     - 新規メモの作成")
	log.Println("  GET    /memos/:id - 特定のメモの取得")
	log.Println("  PUT    /memos/:id - メモの更新")
	log.Println("  DELETE /memos/:id - メモの削除")
	log.Println("サーバーは http://localhost:8080 で実行中です")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("サーバーの起動に失敗しました:", err)
	}
}

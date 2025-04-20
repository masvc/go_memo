package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"memo-api/models"

	"github.com/google/uuid"
)

// memoStore はメモデータをメモリ上で保持する変数です
var memoStore models.MemoStore

// init はパッケージの初期化時に実行されます
// アプリケーション起動時にJSONファイルからデータを読み込みます
func init() {
	loadMemos()
}

// loadMemos はJSONファイルからメモデータを読み込みます
// ファイルが存在しない場合は空のメモリストを作成します
func loadMemos() {
	data, err := ioutil.ReadFile("data/memos.json")
	if err != nil {
		memoStore = models.MemoStore{Memos: []models.Memo{}}
		return
	}
	if err := json.Unmarshal(data, &memoStore); err != nil {
		// JSONのパースに失敗した場合は空のリストを作成
		memoStore = models.MemoStore{Memos: []models.Memo{}}
	}
}

// saveMemos はメモデータをJSONファイルに保存します
func saveMemos() error {
	data, err := json.MarshalIndent(memoStore, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("data/memos.json", data, 0644)
}

// GetAllMemos は全てのメモを取得するハンドラーです
// GET /memos エンドポイントで呼び出されます
func GetAllMemos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(memoStore.Memos); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// GetMemoByID は指定されたIDのメモを取得するハンドラーです
// GET /memos/:id エンドポイントで呼び出されます
func GetMemoByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[len("/memos/"):]
	for _, memo := range memoStore.Memos {
		if memo.ID == id {
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(memo); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
			return
		}
	}
	http.Error(w, "Memo not found", http.StatusNotFound)
}

// CreateMemo は新しいメモを作成するハンドラーです
// POST /memos エンドポイントで呼び出されます
func CreateMemo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var memo models.Memo
	if err := json.NewDecoder(r.Body).Decode(&memo); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 新しいメモの初期化
	memo.ID = uuid.New().String()
	memo.CreatedAt = time.Now()
	memo.UpdatedAt = time.Now()

	memoStore.Memos = append(memoStore.Memos, memo)
	if err := saveMemos(); err != nil {
		http.Error(w, "Failed to save memo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(memo); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// UpdateMemo は既存のメモを更新するハンドラーです
// PUT /memos/:id エンドポイントで呼び出されます
func UpdateMemo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[len("/memos/"):]
	var updatedMemo models.Memo
	if err := json.NewDecoder(r.Body).Decode(&updatedMemo); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, memo := range memoStore.Memos {
		if memo.ID == id {
			updatedMemo.ID = id
			updatedMemo.CreatedAt = memo.CreatedAt
			updatedMemo.UpdatedAt = time.Now()
			memoStore.Memos[i] = updatedMemo
			if err := saveMemos(); err != nil {
				http.Error(w, "Failed to save memo", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(updatedMemo); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
			return
		}
	}
	http.Error(w, "Memo not found", http.StatusNotFound)
}

// DeleteMemo は指定されたIDのメモを削除するハンドラーです
// DELETE /memos/:id エンドポイントで呼び出されます
func DeleteMemo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[len("/memos/"):]
	for i, memo := range memoStore.Memos {
		if memo.ID == id {
			memoStore.Memos = append(memoStore.Memos[:i], memoStore.Memos[i+1:]...)
			if err := saveMemos(); err != nil {
				http.Error(w, "Failed to save changes", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Memo not found", http.StatusNotFound)
}

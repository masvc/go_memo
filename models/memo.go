package models

import (
	"time"
)

// Memo はメモの基本構造を定義します
type Memo struct {
	ID        string    `json:"id"`         // メモの一意の識別子
	Title     string    `json:"title"`      // メモのタイトル
	Content   string    `json:"content"`    // メモの本文
	Tags      []string  `json:"tags"`       // メモに関連付けられたタグのリスト
	CreatedAt time.Time `json:"created_at"` // メモの作成日時
	UpdatedAt time.Time `json:"updated_at"` // メモの最終更新日時
}

// MemoStore はメモの集合を管理する構造体です
// JSONファイルとの相互変換に使用されます
type MemoStore struct {
	Memos []Memo `json:"memos"` // メモのリスト
}

package entity

import "time"

// ユーザーのエンティティ
type User struct {
	ID        uint      // ユーザーID
	Name      string    // ユーザー名
	CreatedAt time.Time // 作成日時
	UpdatedAt time.Time // 更新日時
}

func NewUser(id uint, name string, createdAt time.Time, updatedAt time.Time) *User {
	return &User{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

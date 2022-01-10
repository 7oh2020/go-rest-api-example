package dto

import "time"

// DTOはデータ転送専用の構造体です。
//アプリケーション層とドメイン層の間のやり取りに使用されるほか、リクエストやレスポンスとのバインドにも使用されます。

// IDのみのユーザーデータ
type UserIDModel struct {
	ID uint `query:"id" form:"id" validate:"required"`
}

// ユーザーデータ
type UserModel struct {
	ID        uint      `json:"id" form:"id" validate:"required"`
	Name      string    `json:"name" form:"name" validate:"required,max=32"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserModel(id uint, name string, createdAt time.Time, updatedAt time.Time) *UserModel {
	return &UserModel{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

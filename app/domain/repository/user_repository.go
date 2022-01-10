package repository

import "go-rest-api-example/app/infrastructure/mysql/entity"

// リポジトリのインタフェースを使用するとサービスとの結合度を下げることができます。
// 結合度が下がるとテストが容易になり、追加修正の際の影響範囲も小さく済みます。
type IUserRepository interface {
	FindByID(id uint) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id uint) error
}

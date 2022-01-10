package di

import (
	"database/sql"
	"go-rest-api-example/app/domain/service"
	"go-rest-api-example/app/handler"
	"go-rest-api-example/app/infrastructure/mysql"
)

// DI（依存注入）を使用してハンドラとサービスとリポジトリを合体させています。
// それぞれインタフェースで疎結合されているので、実装を差し替えたり単体テストをするのも容易です。
func InitUser(db *sql.DB) handler.IUserHandler {
	r := mysql.NewUserRepository(db)
	s := service.NewUserService(r)
	return handler.NewUserHandler(s)

}

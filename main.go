package main

import (
	"database/sql"
	"go-rest-api-example/app/common/di"
	"go-rest-api-example/app/common/validation"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &validation.CustomValidator{V: validator.New()}

	// MySQLとの接続を開きます。
	// 接続情報は環境変数MYSQL_DSNから取得しています。
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	setUserRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}

// UserHandlerのメソッドとWEBルートを関連付けします。
func setUserRoutes(e *echo.Echo, db *sql.DB) {
	user := di.InitUser(db)
	e.GET("/user/detail", user.FindByID)
	e.POST("/user/create", user.Create)
	e.POST("/user/update", user.Update)
	e.POST("/user/delete", user.Delete)
}

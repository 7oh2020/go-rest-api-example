package handler

import (
	"go-rest-api-example/app/common"
	"go-rest-api-example/app/common/dto"
	"net/http"
	"time"

	"go-rest-api-example/app/domain/service"

	"github.com/labstack/echo/v4"
)

//インタフェースを使用することで上位との結合度を下げることができます。
// また、インタフェースからモックを自動生成できるためテストが容易になります。
type IUserHandler interface {
	// IDからユーザーを取得します。
	FindByID(c echo.Context) error

	// ユーザーを作成します。
	Create(c echo.Context) error

	// ユーザーを更新します。
	Update(c echo.Context) error

	// ユーザーを削除します
	Delete(c echo.Context) error
}

// ハンドラはアプリケーションレベルのロジックを担います。
// ドメインロジックに関しては内部に埋め込まれたサービスが担います。
type UserHandler struct {
	service.IUserService
}

func NewUserHandler(srv service.IUserService) IUserHandler {
	return &UserHandler{srv}
}
func (h *UserHandler) FindByID(c echo.Context) error {
	var user dto.UserIDModel
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(400, "failed to bind request")
	}
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(400, "failed to validation request")
	}

	ret, err := h.IUserService.FindByID(user.ID)
	if err != nil {
		return echo.NewHTTPError(404, "user is not exists")
	}
	return c.JSON(http.StatusOK, ret)
}

func (h *UserHandler) Create(c echo.Context) error {
	var user dto.UserModel
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(400, "failed to bind request")
	}

	// INSERT時はIDを使用しないため、IDのバリデーションをスキップします。
	user.ID = 1

	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(400, "failed to validation request")
	}
	if now, ok := c.Get("now").(time.Time); ok {
		// ﾃｽﾄ時はコンテキストから時刻を受け取る
		user.CreatedAt = now
	} else {
		// 通常時は現在時刻を取得する
		user.CreatedAt = common.CurrentTime()
	}

	if err := h.IUserService.Create(&user); err != nil {
		return echo.NewHTTPError(500, "failed to create &user")
	}
	return c.String(http.StatusCreated, `{}`)
}

func (h *UserHandler) Update(c echo.Context) error {
	var user dto.UserModel
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(400, err.Error())
		// return echo.NewHTTPError(400, "failed to bind request")
	}
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(400, "failed to validation request")
	}
	if now, ok := c.Get("now").(time.Time); ok {
		// ﾃｽﾄ時はコンテキストから時刻を受け取る
		user.UpdatedAt = now
	} else {
		// 通常時は現在時刻を取得する
		user.UpdatedAt = common.CurrentTime()
	}

	if err := h.IUserService.Update(&user); err != nil {
		return echo.NewHTTPError(500, "failed to create &user")
	}
	return c.String(http.StatusNoContent, `{}`)
}

func (h *UserHandler) Delete(c echo.Context) error {
	var user dto.UserIDModel
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(400, "failed to bind request")
	}
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(400, "failed to validation request")
	}

	if err := h.IUserService.Delete(user.ID); err != nil {
		return echo.NewHTTPError(404, "user is not exists")
	}
	return c.String(http.StatusNoContent, `{}`)
}

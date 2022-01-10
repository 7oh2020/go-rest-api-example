package handler

import (
	"fmt"
	"go-rest-api-example/app/common"
	"go-rest-api-example/app/common/dto"
	"go-rest-api-example/app/common/validation"
	"go-rest-api-example/mocks"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

// testifyパッケージのモック機能を使用してインタフェースからモックを自動生成しています。
// 下位レイヤーをモック化することでアプリケーションロジックのテストに集中できます。
func TestUserHandler_FindByID(tt *testing.T) {
	tt.Run(
		"正常系: エラーなし",
		func(t *testing.T) {
			now := common.CurrentTime()
			user := dto.NewUserModel(1, "Alice", now, now)
			req := httptest.NewRequest(http.MethodGet, "/user/detail?id=1", nil)

			e := echo.New()
			e.Validator = &validation.CustomValidator{V: validator.New()}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			s := new(mocks.IUserService)
			s.On("FindByID", user.ID).Return(user, nil)
			h := NewUserHandler(s)

			if assert.NoError(t, h.FindByID(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				s.AssertExpectations(t)
			}
		})
	tt.Run(
		"準正常系: idに0を指定→バリデーションに失敗",
		func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/user/detail?id=0", nil)

			e := echo.New()
			e.Validator = &validation.CustomValidator{V: validator.New()}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			s := new(mocks.IUserService)
			h := NewUserHandler(s)

			assert.EqualError(t, h.FindByID(c), `code=400, message=failed to validation request`)
		})
	tt.Run(
		"準正常系: idに文字列を指定→バインドに失敗",
		func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/user/detail?id=abc", nil)

			e := echo.New()
			e.Validator = &validation.CustomValidator{V: validator.New()}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			s := new(mocks.IUserService)
			h := NewUserHandler(s)

			assert.EqualError(t, h.FindByID(c), `code=400, message=failed to bind request`)
		})
}

func TestUserHandler_Create(tt *testing.T) {
	tt.Run(
		"正常系: エラーなし",
		func(t *testing.T) {
			now := common.CurrentTime()
			user := &dto.UserModel{ID: 1, Name: "Alice", CreatedAt: now}
			form := make(url.Values)
			form.Set("name", user.Name)
			req := httptest.NewRequest(http.MethodPost, "/user/create", strings.NewReader(form.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

			e := echo.New()
			e.Validator = &validation.CustomValidator{V: validator.New()}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			//テスト用の時刻をコンテキストにセットする
			c.Set("now", now)

			s := new(mocks.IUserService)
			s.On("Create", user).Return(nil)
			h := NewUserHandler(s)

			assert.NoError(t, h.Create(c))
			s.AssertExpectations(t)
		})
	tt.Run(
		"準正常系: nameの長さが最大長(32)+1→バリデーションエラー",
		func(t *testing.T) {
			now := common.CurrentTime()
			user := &dto.UserModel{ID: 1, Name: "0123456789abcdefghijあいうえおかきくけこ＠！？", CreatedAt: now}
			form := make(url.Values)
			form.Set("name", user.Name)
			req := httptest.NewRequest(http.MethodPost, "/user/create", strings.NewReader(form.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

			e := echo.New()
			e.Validator = &validation.CustomValidator{V: validator.New()}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			//テスト用の時刻をコンテキストにセットする
			c.Set("now", now)

			s := new(mocks.IUserService)
			h := NewUserHandler(s)

			assert.EqualError(t, h.Create(c), `code=400, message=failed to validation request`)
		})
}

func TestUserHandler_Update(tt *testing.T) {
	tt.Run(
		"正常系: エラーなし",
		func(t *testing.T) {
			now := common.CurrentTime()
			user := &dto.UserModel{ID: 1, Name: "Alice", UpdatedAt: now}
			form := make(url.Values)
			form.Set("id", fmt.Sprintf("%d", user.ID))
			form.Set("name", user.Name)
			req := httptest.NewRequest(http.MethodPost, "/user/update", strings.NewReader(form.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

			e := echo.New()
			e.Validator = &validation.CustomValidator{V: validator.New()}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			//テスト用の時刻をコンテキストにセットする
			c.Set("now", now)

			s := new(mocks.IUserService)
			s.On("Update", user).Return(nil)
			h := NewUserHandler(s)
			assert.NoError(t, h.Update(c))
			s.AssertExpectations(t)
		})
	tt.Run(
		"準正常系: id指定なし→バリデーションエラー",
		func(t *testing.T) {
			now := common.CurrentTime()
			user := &dto.UserModel{ID: 1, Name: "Alice", UpdatedAt: now}
			form := make(url.Values)
			form.Set("name", user.Name)
			req := httptest.NewRequest(http.MethodPost, "/user/update", strings.NewReader(form.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

			e := echo.New()
			e.Validator = &validation.CustomValidator{V: validator.New()}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			//テスト用の時刻をコンテキストにセットする
			c.Set("now", now)

			s := new(mocks.IUserService)
			h := NewUserHandler(s)

			assert.EqualError(t, h.Update(c), `code=400, message=failed to validation request`)
		})
}

func TestUserHandler_Delete(tt *testing.T) {
	tt.Run(
		"正常系: エラーなし",
		func(t *testing.T) {
			id := 1
			form := make(url.Values)
			form.Set("id", fmt.Sprintf("%d", id))
			req := httptest.NewRequest(http.MethodPost, "/user/delete", strings.NewReader(form.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

			e := echo.New()
			e.Validator = &validation.CustomValidator{V: validator.New()}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			s := new(mocks.IUserService)
			s.On("Delete", uint(id)).Return(nil)
			h := NewUserHandler(s)
			assert.NoError(t, h.Delete(c))
			s.AssertExpectations(t)
		})
}

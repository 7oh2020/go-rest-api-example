package service

import (
	"go-rest-api-example/app/common"
	"go-rest-api-example/app/common/dto"
	"go-rest-api-example/app/infrastructure/mysql/entity"
	"go-rest-api-example/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

// testifyパッケージのモック機能を使用してインタフェースからモックを自動生成しています。
// 下位レイヤーをモック化することでドメインロジックのテストに集中できます。
func TestUserService_FindByID(tt *testing.T) {
	tt.Run(
		"正常系: エラーなし",
		func(t *testing.T) {
			user := &entity.User{
				ID:        1,
				Name:      "Alice",
				CreatedAt: common.CurrentTime(),
				UpdatedAt: common.CurrentTime(),
			}
			r := new(mocks.IUserRepository)
			r.On("FindByID", user.ID).Return(user, nil)

			s := NewUserService(r)
			ret, err := s.FindByID(user.ID)

			assert.NoError(t, err)
			assert.Equal(t, ret.ID, user.ID)
			assert.Equal(t, ret.Name, user.Name)
			assert.Equal(t, ret.CreatedAt, user.CreatedAt)
			assert.Equal(t, ret.UpdatedAt, user.UpdatedAt)
			r.AssertExpectations(t)
		})
}

func TestUserService_Create(tt *testing.T) {
	tt.Run(
		"正常系: エラーなし",
		func(t *testing.T) {
			now := common.CurrentTime()
			user := dto.NewUserModel(1, "Alice", now, now)
			enuser := entity.NewUser(1, "Alice", now, now)

			r := new(mocks.IUserRepository)
			r.On("Create", enuser).Return(nil)

			s := NewUserService(r)
			assert.NoError(t, s.Create(user))
			r.AssertExpectations(t)

		})
}

func TestUserService_Update(tt *testing.T) {
	tt.Run(
		"正常系: エラーなし",
		func(t *testing.T) {
			now := common.CurrentTime()
			user := dto.NewUserModel(1, "Alice", now, now)
			enuser := entity.NewUser(1, "Alice", now, now)

			r := new(mocks.IUserRepository)
			r.On("Update", enuser).Return(nil)

			s := NewUserService(r)
			assert.NoError(t, s.Update(user))
			r.AssertExpectations(t)
		})
}

func TestUserService_Delete(tt *testing.T) {
	tt.Run(
		"正常系: エラーなし",
		func(t *testing.T) {
			id := uint(1)
			r := new(mocks.IUserRepository)
			r.On("Delete", id).Return(nil)

			s := NewUserService(r)
			assert.NoError(t, s.Delete(id))
			r.AssertExpectations(t)
		})
}

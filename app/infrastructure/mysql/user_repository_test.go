package mysql

import (
	"go-rest-api-example/app/common"
	"go-rest-api-example/app/infrastructure/mysql/entity"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

// sqlmockパッケージを使用してデータベースをモック化しています。
// リポジトリ内部で意図したSQLが発行されているか、引数と戻り値が妥当かをチェックします。
func TestUserRepository_Create(tt *testing.T) {
	tt.Run(
		"正常系: エラーなし",
		func(t *testing.T) {
			user := &entity.User{
				Name:      "Alice",
				CreatedAt: common.CurrentTime(),
			}
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(`INSERT INTO users`).
				ExpectExec().
				WithArgs(user.Name, user.CreatedAt).
				WillReturnResult(sqlmock.NewResult(1, 1))

			r := NewUserRepository(db)
			assert.NoError(t, r.Create(user))
			assert.NoError(t, mock.ExpectationsWereMet())
		})
}

func TestUserRepository_Update(tt *testing.T) {
	tt.Run(
		"正常系: エラーなし",
		func(t *testing.T) {
			user := &entity.User{
				ID:        1,
				Name:      "Alice",
				UpdatedAt: common.CurrentTime(),
			}
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(`UPDATE users`).
				ExpectExec().
				WithArgs(user.Name, user.UpdatedAt, user.ID).
				WillReturnResult(sqlmock.NewResult(1, 1))

			r := NewUserRepository(db)
			assert.NoError(t, r.Update(user))
			assert.NoError(t, mock.ExpectationsWereMet())
		})
}

func TestUserRepository_Delete(tt *testing.T) {
	tt.Run(
		"正常系: エラーなし",
		func(t *testing.T) {
			id := uint(1)
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(`DELETE FROM users`).
				ExpectExec().
				WithArgs(id).
				WillReturnResult(sqlmock.NewResult(1, 1))

			r := NewUserRepository(db)
			assert.NoError(t, r.Delete(id))
			assert.NoError(t, mock.ExpectationsWereMet())
		})
}

func TestUserRepository_FindByID(tt *testing.T) {
	tt.Run(
		"正常系: エラーなし",
		func(t *testing.T) {
			user := &entity.User{
				ID:        1,
				Name:      "Alice",
				CreatedAt: common.CurrentTime(),
				UpdatedAt: common.CurrentTime(),
			}
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(`SELECT id, name, created_at, updated_at FROM users`).
				ExpectQuery().
				WithArgs(user.ID).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).AddRow(user.ID, user.Name, user.CreatedAt, user.UpdatedAt))

			r := NewUserRepository(db)
			ret, err := r.FindByID(user.ID)

			assert.NoError(t, err)
			assert.Equal(t, ret.ID, user.ID)
			assert.Equal(t, ret.Name, user.Name)
			assert.Equal(t, ret.CreatedAt, user.CreatedAt)
			assert.Equal(t, ret.UpdatedAt, user.UpdatedAt)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
}

package mysql

import (
	"database/sql"
	"go-rest-api-example/app/domain/repository"
	"go-rest-api-example/app/infrastructure/mysql/entity"
)

//リポジトリはデータベースへのコネクションを持ち、エンティティを永続化します。
// 永続化の方法はリポジトリ内部にカプセル化されます。
type UserRepository struct {
	*sql.DB
}

func NewUserRepository(db *sql.DB) repository.IUserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindByID(id uint) (*entity.User, error) {
	stmt, err := r.Prepare(`SELECT id, name, created_at, updated_at FROM users WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &entity.User{}
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Create(user *entity.User) error {
	stmt, err := r.Prepare(`INSERT INTO users(name, created_at) VALUES (?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Name, user.CreatedAt); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(user *entity.User) error {
	stmt, err := r.Prepare(`UPDATE users SET name = ?, updated_at = ? WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Name, user.UpdatedAt, user.ID); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id uint) error {
	stmt, err := r.Prepare(`DELETE FROM users WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil

}

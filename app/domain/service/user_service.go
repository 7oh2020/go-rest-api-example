package service

import (
	"go-rest-api-example/app/common/dto"
	"go-rest-api-example/app/domain/repository"
	"go-rest-api-example/app/infrastructure/mysql/entity"
)

// インタフェースを使用することでハンドラとの結合度を下げます。
// また、インタフェースからモックを自動生成できるためテストが容易になります。
type IUserService interface {
	FindByID(id uint) (*dto.UserModel, error)
	Create(user *dto.UserModel) error
	Update(user *dto.UserModel) error
	Delete(id uint) error
}

// サービスはドメインレベルのロジックを担います。
//データの永続化は内部に埋め込まれたリポジトリが担います。
// また、エンティティをDTOに変換してドメイン知識がアプリケーション層に流出するのを防ぎます。
type UserService struct {
	repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{repo}
}

// エンティティからDTOに変換する
func (s *UserService) convertTo(user *entity.User) *dto.UserModel {
	return dto.NewUserModel(user.ID, user.Name, user.CreatedAt, user.UpdatedAt)
}

// DTOからエンティティに変換する
func (s *UserService) convertFrom(user *dto.UserModel) *entity.User {
	return entity.NewUser(user.ID, user.Name, user.CreatedAt, user.UpdatedAt)
}

func (s *UserService) FindByID(id uint) (*dto.UserModel, error) {
	user, err := s.IUserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return s.convertTo(user), nil
}

func (s *UserService) Create(user *dto.UserModel) error {
	u := s.convertFrom(user)
	return s.IUserRepository.Create(u)
}

func (s *UserService) Update(user *dto.UserModel) error {
	u := s.convertFrom(user)
	return s.IUserRepository.Update(u)
}

func (s *UserService) Delete(id uint) error {
	return s.IUserRepository.Delete(id)
}

package repository

import (
	"echo-rest-api-practice/entities"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *entities.User, email string) error
	CreateUser(user *entities.User) error
}

type userRepository struct {
	db *gorm.DB
}

// userRepositoryを生成する関数
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// DBからユーザの取得を行う関数
func (ur *userRepository) GetUserByEmail(user *entities.User, email string) error {
	// 引数で受け取ったemailを持つUserをDBから取得し、userにセットする
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// DBへユーザの登録を行う関数
func (ur *userRepository) CreateUser(user *entities.User) error {
	// 引数で受け取ったuserをDBに登録し、userを登録した情報で書き換える
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

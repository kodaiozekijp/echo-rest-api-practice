package repository

import "echo-rest-api-practice/entities"

type IUserRepository interface {
	GetUserByEmail(user *entities.User, email string) error
	CreateUser(user *entities.User) error
}


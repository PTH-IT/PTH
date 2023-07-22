package repository

import (
	"Golang/domain/model"
)

type UserRepository interface {
	GetUser(userId string, password string) (*model.User, error)
	AddUser(userId string, password string) error
}

package repository

import (
	"Golang/domain/model"
)

type UserRepository interface {
	GetUser(userName string, password string) (*model.User, error)
	AddUser(userName string, password string, email string) (*model.User, error)
}

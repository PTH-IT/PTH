package usecase

import (
	"fmt"

	"Golang/domain/model"
	"Golang/domain/repository"
	InforLog "Golang/log/infor"
)

type Reference interface {
	GetUser(userName string, password string) (*model.User, error)
	AddUser(userName string, password string, email string) (*model.User, error)
	UpdateConnectionID(userId string, connectionid string) error
	GetConnectionID(userId string) (*model.GetUser, error)
}

func NewReferrance(
	userRepository repository.UserRepository,
) Reference {
	return reference{
		userRepository,
	}
}

type reference struct {
	userRepository repository.UserRepository
}

func (r reference) GetUser(userId string, password string) (*model.User, error) {

	user, err := r.userRepository.GetUser(userId, password)
	return user, err
}
func (r reference) AddUser(userName string, password string, email string) (*model.User, error) {

	user, err := r.userRepository.AddUser(userName, password, email)
	return user, err
}
func (r reference) GetConnectionID(userId string) (*model.GetUser, error) {
	InforLog.PrintLog(fmt.Sprintf("r.mongoRepository.GetUser call"))
	// user, err := r.mongoRepository.GetConnectionID(userId)
	// return user, err
	return nil, nil
}
func (r reference) UpdateConnectionID(userId string, connectionid string) error {
	InforLog.PrintLog(fmt.Sprintf("r.mongoRepository.GetUser call"))
	// err := r.mongoRepository.UpdateConnectionID(userId, connectionid)
	// return err
	return nil

}

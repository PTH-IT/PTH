package gormdb

import (
	"Golang/domain/model"
	"Golang/domain/repository"
	"Golang/utils"

	"gorm.io/gorm"
)

func NewUser() repository.UserRepository {
	return userRepository{}
}

type userRepository struct {
}

func (repo userRepository) GetUser(userName string, password string) (*model.User, error) {
	var user []*model.User
	DB.Table("tb_user").Where("user_name  = ? and password = ?", userName, password).Find(&user)
	if len(user) == 0 {
		return nil, nil
	}
	return user[0], nil
}
func (repo userRepository) AddUser(userName string, password string, email string) (*model.User, error) {
	var err error
	user := &model.User{
		UserName:    userName,
		Password:    password,
		Email:       email,
		Status:      "0",
		CreatedTime: utils.GettimeNow(),
		UpdatedTime: utils.GettimeNow(),
	}
	var existingUser *model.User
	if err = DB.Table("tb_user").Where("email = ? or user_name = ? ", user.Email, user.UserName).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Record not found, proceed with insertion
			err = DB.Table("tb_user").Create(user).Error
			if err != nil {
				return nil, err
			}
		} else {
			return existingUser, nil
		}
	}

	return nil, nil
}

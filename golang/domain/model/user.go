package model

type User struct {
	ID          int64  `gorm:"column:id"`
	UserName    string `gorm:"column:user_name"`
	Password    string `gorm:"column:password"`
	Email       string `gorm:"column:email"`
	Status      string `gorm:"column:status"`
	CreatedTime string `gorm:"column:created_time"`
	UpdatedTime string `gorm:"column:updated_time"`
}
type Login struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type RegisterUser struct {
	ID          int64  `gorm:"column:id"`
	UserName    string `gorm:"column:user_name"`
	Password    string `gorm:"column:password"`
	Email       string `gorm:"column:email"`
	Status      string `gorm:"column:status"`
	CreatedTime string `gorm:"column:created_time"`
	UpdatedTime string `gorm:"column:updated_time"`
}

type GetUser struct {
	UserName     string `bson:"userid"`
	Password     string `bson:"password"`
	Email        string `bson:"email"`
	ActiceEmail  string `bson:"actice"`
	Status       string `bson:"status"`
	CreatedTime  string `bson:"createdtime"`
	UpdatedTime  string `bson:"updatedtime"`
	ConnectionId string `bson:"connectionid"`
}

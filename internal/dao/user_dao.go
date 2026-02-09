package dao

import (
	"go-web-starter/internal/model"
	"go-web-starter/pkg/db"
)

type UserCreateReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetByUsername(name string) (*model.User, error) {
	var u model.User
	err := db.DB.Where("username = ?", name).First(&u).Error
	return &u, err
}

func List() []model.User {
	var users []model.User
	db.DB.Find(&users)
	return users
}

func Create(req UserCreateReq) {
	db.DB.Create(&model.User{
		Username: req.Username,
		Password: req.Password,
	})
}

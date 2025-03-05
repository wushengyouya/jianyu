package model

import (
	"errors"

	"gorm.io/gorm"
)

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ? is_del = ?", a.AppKey, a.AppSecret, 0)
	err := db.First(&auth).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return auth, err
	}
	return auth, nil
}

func (a Auth) TableName() string {
	return "blog_auth"
}

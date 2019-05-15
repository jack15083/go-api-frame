package models

import (
	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID        int
	Username  string
	Password  string
	Avatar    string
	Mobile    string
	LastLogin float64
	LastIp    string
	TryTime   int
}

func (UserModel) TableName() string {
	return "yy_auth_admin"
}

func (u *UserModel) CheckPassword(plainPwd string) bool {
	byteHash := []byte(u.Password)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		return false
	}
	return true
}

func (u *UserModel) GetFirstByUsername(db *gorm.DB, username string) {
	db.Where("username = ?", username).First(&u)
	return
}

package model

import "time"

type User struct {
	UserId     int       `gorm:"primary_key" json:"user_id"`
	Username   string    `json:"username"`
	Nickname   string    `json:"nickname"`
	Password   string    `json:"password"`
	Salt       string    `json:"salt"`
	AddDate    time.Time `json:"add_date" gorm:"-"`
	ChangeDate time.Time `json:"change_date" gorm:"-"`
	IsDelete   int       `json:"is_delete"`
}

func (User) TableName() string {
	return "blog_user"
}

func GetUserInfoByUsername(username string, userInfo *User) error {
	return db.Model(&User{}).Where("username = ?", username).First(userInfo).Error
}

package model

import "time"

type Token struct {
	TId         int       `gorm:"primary_key" json:"t_id"`
	UserId      int       `json:"user_id"`
	AccessToken string    `json:"access_token"`
	Ip          string    `json:"ip"`
	ExpireTime  int       `json:"expire_time"`
	AddTime     int       `json:"add_time"`
	AddDate     time.Time `json:"add_date" gorm:"-"`
	IsDelete    int       `json:"is_delete"`
	UpdateDate  time.Time `json:"update_date" gorm:"-"`
}

func (Token) TableName() string {
	return "lt_user_token"
}

func CreateUserToken(userId int, accessToken string, ip string, expireTime int) (err error) {
	token := Token{
		UserId:      userId,
		AccessToken: accessToken,
		Ip:          ip,
		ExpireTime:  expireTime,
	}
	return db.Create(&token).Error
}

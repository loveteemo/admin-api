package model

import "time"

type Token struct {
	LogId       int       `gorm:"primary_key" json:"log_id"`
	UserId      int       `json:"user_id"`
	AccessToken string    `json:"access_token"`
	Ip          string    `json:"ip"`
	ExpireTime  int       `json:"expire_time"`
	IsDelete    int       `json:"is_delete"`
	AddDate     time.Time `json:"add_date" gorm:"-"`
	UpdateDate  time.Time `json:"update_date" gorm:"-"`
}

func (Token) TableName() string {
	return "blog_user_token"
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

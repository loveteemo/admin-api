package model

type User struct {
	UserId   int    `gorm:"primary_key" json:"user_id"`
	Mobile   string `json:"mobile"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	AddTime  int    `json:"add_time"`
	AddDate  string `json:"add_date"`
	IsDelete int    `json:"is_delete"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

func (User) TableName() string {
	return "lt_user"
}

func GetUserInfoByUsername(username string, userInfo *User) error {
	return db.Model(&User{}).Where("username = ?", username).First(userInfo).Error
}

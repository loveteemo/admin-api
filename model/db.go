package model

import (
	"blog-admin-api/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func Initialization() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPort, utils.DbName)
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println("conn db err")
	}

	//调试模式
	if utils.AppMode == "debug" {
		db = db.Debug()
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println("conn db err")
	}

	sqlDb.SetMaxIdleConns(10)

	sqlDb.SetMaxOpenConns(100)

	sqlDb.SetConnMaxLifetime(10 * time.Second)

}

func Close() {
	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println("conn db err")
	}
	_ = sqlDb.Close()
}

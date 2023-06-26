package main

import (
	"blog-admin-api/model"
	"blog-admin-api/router"
)

func main() {
	//初始化数据库
	model.Initialization()

	//关闭数据库
	defer func() {
		model.Close()
	}()

	//添加路由
	router.Init()
}

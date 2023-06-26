package router

import (
	"blog-admin-api/controller"
	"blog-admin-api/middleware"
	"blog-admin-api/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {

	//设置gin模式
	gin.SetMode(utils.AppMode)

	router := gin.Default()

	router.Use(cors.Default())

	router.Group("/").Use(middleware.Request())
	{
		//登录
		router.POST("/login", controller.Login)
	}

	err := router.Run(utils.HttpPort)
	if err != nil {
		return
	}
}

package router

import (
	adminController "blog-admin-api/blog/admin/controller"
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

	// 后台API
	admin := router.Group("/blog/admin").Use(middleware.Request())
	{
		admin.POST("/login", adminController.Login)
	}

	// 前台API
	router.Group("/blog/api")
	{

	}

	err := router.Run(utils.HttpPort)
	if err != nil {
		return
	}
}

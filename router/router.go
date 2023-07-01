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

	//创建路由
	router := gin.Default()

	// 跨域
	router.Use(cors.Default())

	// 后台API
	admin := router.Group("/blog/admin").Use(middleware.Request())
	{
		admin.POST("/login", adminController.Login)
		admin.POST("/article/list", adminController.ArticleList)
		admin.POST("/article/edit", adminController.ArticleEdit)
		admin.POST("/article/detail", adminController.ArticleDetail)
		admin.POST("/article/delete", adminController.ArticleDelete)
		admin.POST("category/list", adminController.CategoryList)
		admin.POST("category/edit", adminController.CategoryEdit)
		admin.POST("category/delete", adminController.CategoryDelete)
		admin.POST("category/detail", adminController.CategoryDetail)
	}

	// 前台API
	router.Group("/blog/api")
	{

	}

	//启动路由
	err := router.Run(utils.HttpPort)
	if err != nil {
		return
	}
}

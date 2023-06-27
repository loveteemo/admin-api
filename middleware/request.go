package middleware

import (
	"blog-admin-api/utils"
	"github.com/gin-gonic/gin"
)

func Request() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		blogToken := ctx.GetHeader("Authorization")
		if blogToken == "" && ctx.Request.URL.Path != "/blog/admin/login" {
			utils.ReturnResult(ctx, 10000, "请登录", nil)
			return
		}

		ctx.Next()
	}
}

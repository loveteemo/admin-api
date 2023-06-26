package controller

import (
	"blog-admin-api/model"
	"blog-admin-api/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(ctx *gin.Context) {
	//请求参数
	var request LoginRequest

	//绑定参数到request
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		utils.ReturnResult(ctx, -1, "参数错误", nil)
		return
	}

	if request.Username == "" {
		utils.ReturnResult(ctx, -1, "用户名不能为空", nil)
		return
	}
	if request.Password == "" {
		utils.ReturnResult(ctx, -1, "密码不能为空", nil)
		return
	}

	//查询密码
	var userInfo model.User
	if err = model.GetUserInfoByUsername(request.Username, &userInfo); err != nil {
		utils.ReturnResult(ctx, -1, "用户不存在", nil)
		return
	}
	if utils.EncryptPassword(request.Password, userInfo.Salt) != userInfo.Password {
		utils.ReturnResult(ctx, -1, "密码错误", nil)
		return
	}

	tokenString := utils.GenerateToken()
	response := LoginResponse{
		Token: tokenString,
	}
	_ = model.CreateUserToken(userInfo.UserId, tokenString, ctx.ClientIP(), int(time.Now().Unix()+3600*24*7))
	//返回参数
	utils.ReturnResult(ctx, 0, "success", response)
	return
}

// ===  参数部分

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Token string `json:"token"`
}

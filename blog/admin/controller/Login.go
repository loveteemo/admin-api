package controller

import (
	"blog-admin-api/model"
	"blog-admin-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(ctx *gin.Context) {
	//请求参数
	var request LoginRequest

	//绑定参数到request
	err := ctx.ShouldBind(&request)
	fmt.Printf("%+v", request)
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
		User: LoginResponseUser{
			Username: userInfo.Username,
			Nickname: userInfo.Nickname,
		},
	}
	_ = model.CreateUserToken(userInfo.UserId, tokenString, ctx.ClientIP(), int(time.Now().Unix()+3600*24*7))
	//返回参数
	utils.ReturnResult(ctx, 0, fmt.Sprintf("登录完成，欢迎回来 %s", userInfo.Nickname), response)
	return
}

// ===  参数部分

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string            `json:"token"`
	User  LoginResponseUser `json:"user"`
}

type LoginResponseUser struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

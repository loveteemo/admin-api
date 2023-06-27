package controller

import (
	"blog-admin-api/model"
	"blog-admin-api/utils"
	"github.com/gin-gonic/gin"
)

func ArticleList(ctx *gin.Context) {
	request := ArticleListRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ReturnResult(ctx, -1, "参数错误", nil)
		return
	}

	var list []model.Article
	if err := model.GetArticleList(request.Page, request.Status, request.Keyword, &list); err != nil {
		utils.ReturnResult(ctx, -1, "获取失败", nil)
		return
	}

	utils.ReturnResult(ctx, 0, "success", list)
	return
}

func ArticleDetail() {

}

func ArticleEdit() {

}

func ArticleDelete() {

}

// ====  参数部分

type ArticleListRequest struct {
	Page    int
	Status  int
	Keyword string
}

type ArticleListResponse struct {
	Total int
	List  []ArticleListResponseItem
}

type ArticleListResponseItem struct {
	ArticleId    int
	Title        string
	Thumbnail    string
	Describe     string
	Hit          int
	CommentTotal int
	CategoryName string
}

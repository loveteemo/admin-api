package controller

import (
	"blog-admin-api/model"
	"blog-admin-api/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func ArticleList(ctx *gin.Context) {
	request := BlogArticleListRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ReturnResult(ctx, -1, "参数错误", nil)
		return
	}

	repo := model.NewBlogArticleRepo()
	var list []model.BlogArticleListItem
	if err := repo.GetList(request.Page, request.Status, request.Keyword, &list); err != nil {
		utils.ReturnResult(ctx, -1, "获取列表失败"+err.Error(), nil)
		return
	}

	var total int64
	if err := repo.GetTotal(request.Status, request.Keyword, &total); err != nil {
		utils.ReturnResult(ctx, -1, "获取total失败"+err.Error(), nil)
		return
	}

	response := BlogArticleListResponse{}
	for _, article := range list {
		response.List = append(response.List, BlogArticleListResponseItem{
			ArticleId:    article.ArticleId,
			Title:        article.Title,
			Thumbnail:    article.Thumbnail,
			Describe:     article.Describe,
			CommentTotal: article.CommentTotal,
			CategoryName: article.CategoryName,
			AddDate:      article.AddDate.Format("2006-01-02 15:04:05"),
			ChangeDate:   article.ChangeDate.Format("2006-01-02 15:04:05"),
		})
	}
	response.Total = total
	utils.ReturnResult(ctx, 0, "success", response)
	return
}

func ArticleEdit(ctx *gin.Context) {
	request := BlogArticleEditRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ReturnResult(ctx, -1, "参数错误", nil)
		return
	}

	detail := model.BlogArticle{
		Title:      request.Title,
		Thumbnail:  request.Thumbnail,
		Describe:   request.Describe,
		Tags:       request.Tags,
		CategoryId: request.CategoryId,
		Content:    request.Content,
		Status:     request.Status,
		Url:        request.Url,
		IsOriginal: request.IsOriginal,
		Author:     request.Author,
		Address:    utils.GetAddress(ctx.ClientIP()),
	}

	repo := model.NewBlogArticleRepo()
	if request.ArticleId == 0 {
		detail.AddDate = time.Now()
		if err := repo.AddDetail(&detail); err != nil {
			utils.ReturnResult(ctx, -1, "添加文章失败"+err.Error(), nil)
			return
		}
		utils.ReturnResult(ctx, 0, "添加文章完成", nil)
		return
	}
	detail.ArticleId = request.ArticleId
	detail.ChangeDate = time.Now()
	if err := repo.UpdateDetail(detail); err != nil {
		utils.ReturnResult(ctx, -1, "修改文章失败"+err.Error(), nil)
		return
	}
	utils.ReturnResult(ctx, 0, "修改文章完成", nil)
	return
}

func ArticleDetail(ctx *gin.Context) {
	request := BlogArticleDetailRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ReturnResult(ctx, -1, "参数错误", nil)
		return
	}
	if request.ArticleId == 0 {
		utils.ReturnResult(ctx, -1, "文章ID错误", nil)
		return
	}

	detail := model.BlogArticle{}
	repo := model.NewBlogArticleRepo()
	if err := repo.Detail(request.ArticleId, &detail); err != nil {
		utils.ReturnResult(ctx, -1, "获取文章详情失败"+err.Error(), nil)
		return
	}
	response := BlogArticleDetailResponse{
		ArticleId:   detail.ArticleId,
		Title:       detail.Title,
		Thumbnail:   detail.Thumbnail,
		Describe:    detail.Describe,
		Tags:        detail.Tags,
		CategoryId:  detail.CategoryId,
		Content:     detail.Content,
		Status:      detail.Status,
		AccessTotal: detail.AccessTotal,
		Url:         detail.Url,
		IsOriginal:  detail.IsOriginal,
		Author:      detail.Author,
		AddDate:     detail.AddDate.Format("2006-01-02 15:04:05"),
	}

	utils.ReturnResult(ctx, 0, "success", response)
	return
}

func ArticleDelete(ctx *gin.Context) {
	request := BlogArticleDeleteRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ReturnResult(ctx, -1, "参数错误", nil)
		return
	}
	if request.ArticleId == 0 {
		utils.ReturnResult(ctx, -1, "文章ID错误", nil)
	}
	repo := model.NewBlogArticleRepo()
	if err := repo.Delete(request.ArticleId); err != nil {
		utils.ReturnResult(ctx, -1, "删除文章失败"+err.Error(), nil)
		return
	}
	utils.ReturnResult(ctx, 0, "删除文章完成", nil)
	return
}

// ====  参数部分

type BlogArticleDetailResponse struct {
	ArticleId    int    `json:"article_id"`
	Title        string `json:"title"`
	Thumbnail    string `json:"thumbnail"`
	Describe     string `json:"describe"`
	Tags         string `json:"tags"`
	CategoryId   int    `json:"category_id"`
	Content      string `json:"content"`
	Status       int    `json:"status"`
	CommentTotal int    `json:"comment_total"`
	AccessTotal  int    `json:"access_total"`
	Url          string `json:"url"`
	IsOriginal   int    `json:"is_original"`
	Author       string `json:"author"`
	Address      string `json:"address"`
	AddDate      string `json:"add_date"`
}

type ArticleIdRequest struct {
	ArticleId int `json:"article_id"`
}

type BlogArticleDeleteRequest struct {
	ArticleIdRequest
}

type BlogArticleDetailRequest struct {
	ArticleIdRequest
}

type BlogArticleEditRequest struct {
	model.BlogArticle
}

type BlogArticleListRequest struct {
	Page    int
	Status  int
	Keyword string
}

type BlogArticleListResponse struct {
	Total int64                         `json:"total"`
	List  []BlogArticleListResponseItem `json:"list"`
}

type BlogArticleListResponseItem struct {
	ArticleId    int    `json:"article_id"`
	Title        string `json:"title"`
	Thumbnail    string `json:"thumbnail"`
	Describe     string `json:"describe"`
	CommentTotal int    `json:"comment_total"`
	CategoryName string `json:"category_name"`
	AddDate      string `json:"add_date"`
	ChangeDate   string `json:"change_date"`
}

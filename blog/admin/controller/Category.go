package controller

import (
	"blog-admin-api/model"
	"blog-admin-api/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func CategoryList(ctx *gin.Context) {

	repo := model.NewBlogCategoryRepo()
	var list []model.BlogCategoryListItem
	if err := repo.GetList(&list); err != nil {
		utils.ReturnResult(ctx, -1, "获取列表失败"+err.Error(), nil)
		return
	}

	response := BlogCategoryListResponse{}
	categoryList := make([]BlogCategoryListResponseItem, 0)
	for _, item := range list {
		categoryList = append(categoryList, BlogCategoryListResponseItem{
			CategoryId: item.CategoryId,
			Title:      item.Title,
			Describe:   item.Describe,
			ParentId:   item.ParentId,
			Url:        item.Url,
			Sort:       item.Sort,
			IsShow:     item.IsShow,
			AddDate:    item.AddDate.Format("2006-01-02 15:04:05"),
			ChangeDate: item.ChangeDate.Format("2006-01-02 15:04:05"),
			Children:   nil,
		})
	}

	tree, err := buildMenuTree(categoryList, 0)
	if err != nil {
		return
	}
	response.List = tree
	utils.ReturnResult(ctx, 0, "success", response)
	return
}

func CategoryEdit(ctx *gin.Context) {
	request := CategoryEditRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ReturnResult(ctx, -1, "参数错误", nil)
		return
	}

	detail := model.BlogCategory{
		Title:    request.Title,
		Describe: request.Describe,
		ParentId: request.ParentId,
		Url:      request.Url,
		Sort:     request.Sort,
		IsShow:   request.IsShow,
	}

	repo := model.NewBlogCategoryRepo()
	if request.CategoryId == 0 {
		detail.AddDate = time.Now()
		if err := repo.AddDetail(&detail); err != nil {
			utils.ReturnResult(ctx, -1, "添加栏目失败"+err.Error(), nil)
			return
		}
		utils.ReturnResult(ctx, 0, "添加栏目完成", nil)
		return
	}
	detail.CategoryId = request.CategoryId
	detail.ChangeDate = time.Now()
	if err := repo.UpdateDetail(detail); err != nil {
		utils.ReturnResult(ctx, -1, "修改栏目失败"+err.Error(), nil)
		return
	}
	utils.ReturnResult(ctx, 0, "修改栏目完成", nil)
	return
}

func CategoryDetail(ctx *gin.Context) {
	request := CategoryDetailRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ReturnResult(ctx, -1, "参数错误", nil)
		return
	}
	if request.CategoryId == 0 {
		utils.ReturnResult(ctx, -1, "栏目ID错误", nil)
		return
	}

	detail := model.BlogCategory{}
	repo := model.NewBlogCategoryRepo()
	if err := repo.Detail(request.CategoryId, &detail); err != nil {
		utils.ReturnResult(ctx, -1, "获取栏目详情失败"+err.Error(), nil)
		return
	}

	response := CategoryDetailResponse{
		CategoryId: detail.CategoryId,
		Title:      detail.Title,
		Describe:   detail.Describe,
		ParentId:   detail.ParentId,
		Url:        detail.Url,
		Sort:       detail.Sort,
		IsShow:     detail.IsShow,
	}
	utils.ReturnResult(ctx, 0, "success", response)
	return
}

func CategoryDelete(ctx *gin.Context) {
	request := CategoryDeleteRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ReturnResult(ctx, -1, "参数错误", nil)
		return
	}
	if request.CategoryId == 0 {
		utils.ReturnResult(ctx, -1, "栏目ID错误", nil)
	}
	repo := model.NewBlogCategoryRepo()
	if err := repo.Delete(request.CategoryId); err != nil {
		utils.ReturnResult(ctx, -1, "删除栏目失败"+err.Error(), nil)
		return
	}
	utils.ReturnResult(ctx, 0, "删除栏目完成", nil)
	return
}

// ====  部分函数
func buildMenuTree(list []BlogCategoryListResponseItem, parentId int) ([]BlogCategoryListResponseItem, error) {
	var tree []BlogCategoryListResponseItem
	for _, item := range list {
		if item.ParentId == parentId {
			children, err := buildMenuTree(list, item.CategoryId)
			if err != nil {
				return nil, err
			}
			if children == nil {
				children = make([]BlogCategoryListResponseItem, 0)
			}
			item.Children = children
			tree = append(tree, item)
		}
	}
	return tree, nil
}

// ====  参数部分

type CategoryDetailResponse struct {
	CategoryId int    `json:"category_id"`
	Title      string `json:"title"`
	Describe   string `json:"describe"`
	ParentId   int    `json:"parent_id"`
	Url        string `json:"url"`
	Sort       int    `json:"sort"`
	IsShow     int    `json:"is_show"`
}

type CategoryIdRequest struct {
	CategoryId int `json:"category_id"`
}

type CategoryDeleteRequest struct {
	CategoryIdRequest
}

type CategoryDetailRequest struct {
	CategoryIdRequest
}

type CategoryEditRequest struct {
	model.BlogCategory
}

type BlogCategoryListResponse struct {
	List []BlogCategoryListResponseItem `json:"list"`
}

type BlogCategoryListResponseItem struct {
	CategoryId int                            `json:"category_id" description:"栏目id"`
	Title      string                         `json:"title" description:"栏目名称"`
	Describe   string                         `json:"describe" description:"栏目描述"`
	ParentId   int                            `json:"parent_id" description:"父级栏目id"`
	Url        string                         `json:"url" description:"栏目链接"`
	Sort       int                            `json:"sort" description:"排序"`
	IsShow     int                            `json:"is_show" description:"是否显示"`
	AddDate    string                         `json:"add_date" gorm:"-" description:"添加时间"`
	ChangeDate string                         `json:"change_date" gorm:"-" description:"修改时间"`
	Children   []BlogCategoryListResponseItem `json:"children" description:"子栏目"`
}

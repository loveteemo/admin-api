package model

import (
	"gorm.io/gorm"
	"time"
)

type BlogCategoryRepo struct {
	db *gorm.DB
}

func NewBlogCategoryRepo() *BlogCategoryRepo {
	return &BlogCategoryRepo{db: db.Table("blog_category")}
}

func (r *BlogCategoryRepo) Delete(categoryId int) error {
	return r.db.Where("category_id = ?", categoryId).Update("is_delete", 1).Error
}

func (r *BlogCategoryRepo) Detail(categoryId int, info *BlogCategory) (err error) {
	return r.db.Where("category_id = ?", categoryId).First(info).Error
}

func (r *BlogCategoryRepo) AddDetail(blogCategory *BlogCategory) (err error) {
	return r.db.Create(blogCategory).Error
}

func (r *BlogCategoryRepo) UpdateDetail(blogCategory BlogCategory) (err error) {
	return r.db.Where("category_id = ?", blogCategory.CategoryId).Updates(blogCategory).Error
}

func (r *BlogCategoryRepo) GetList(list *[]BlogCategoryListItem) (err error) {

	sql := "SELECT a.*, b.total as article_total FROM blog_category a " +
		"LEFT JOIN (SELECT COUNT(1) AS total,category_id FROM blog_article GROUP BY category_id) b ON a.category_id = b.category_id " +
		"WHERE 1=1 "
	sql += " ORDER BY a.parent_id ASC "

	return r.db.Raw(sql).Scan(list).Error
}

func (r *BlogCategoryRepo) GetTotal(keyword string, total *int64) (err error) {

	sql := "SELECT count(1) FROM blog_category a " +
		"LEFT JOIN (SELECT COUNT(1) AS total,category_id FROM blog_article GROUP BY category_id) b ON a.category_id = b.category_id " +
		"WHERE 1=1 "
	var where []interface{}
	if keyword != "" {
		sql += " AND title like '?'"
		where = append(where, "%"+keyword+"%")
	}
	return r.db.Raw(sql, where...).Count(total).Error
}

type BlogCategory struct {
	CategoryId int       `json:"category_id" gorm:"column:category_id;primary_key" description:"栏目id"`
	Title      string    `json:"title" description:"栏目名称"`
	Describe   string    `json:"describe" description:"栏目描述"`
	ParentId   int       `json:"parent_id" description:"父级栏目id"`
	Url        string    `json:"url" description:"栏目链接"`
	Sort       int       `json:"sort" description:"排序"`
	IsShow     int       `json:"is_show" description:"是否显示"`
	AddDate    time.Time `json:"add_date" description:"添加时间"`
	ChangeDate time.Time `json:"change_date" description:"修改时间"`
	IsDelete   int       `json:"is_delete" description:"是否删除"`
}

type BlogCategoryListItem struct {
	BlogCategory
	ArticleTotal int `json:"article_total" description:"文章数量"`
}

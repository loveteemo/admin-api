package model

import (
	"gorm.io/gorm"
	"time"
)

type BlogArticleRepo struct {
	db *gorm.DB
}

// NewBlogArticleRepo 文章模型
// 类似 类的实例化
func NewBlogArticleRepo() *BlogArticleRepo {
	return &BlogArticleRepo{db: db.Table("blog_article")}
}

func (r *BlogArticleRepo) Delete(articleId int) error {
	return r.db.Where("article_id = ?", articleId).Update("is_delete", 1).Error
}

func (r *BlogArticleRepo) Detail(articleId int, article *BlogArticle) (err error) {
	return r.db.Where("article_id = ?", articleId).First(article).Error
}

func (r *BlogArticleRepo) AddDetail(blogArticle *BlogArticle) (err error) {
	return r.db.Create(blogArticle).Error
}

func (r *BlogArticleRepo) UpdateDetail(article BlogArticle) (err error) {
	return r.db.Where("article_id = ?", article.ArticleId).Updates(article).Error
}

func (r *BlogArticleRepo) GetList(page int, status int, keyword string, list *[]BlogArticleListItem) (err error) {

	sql := "SELECT a.*, c.title as category_name FROM blog_article a " +
		"JOIN blog_category c ON a.category_id = c.category_id " +
		"WHERE 1=1 "
	var where []interface{}
	if status != 0 {
		sql += " AND a.status = ? "
		where = append(where, status)
	}
	if keyword != "" {
		sql += " AND title like '?'"
		where = append(where, "%"+keyword+"%")
	}
	sql += " ORDER BY a.article_id DESC "
	sql += " LIMIT ?, 10"
	where = append(where, (page-1)*10)

	return r.db.Raw(sql, where...).Scan(list).Error
}

func (r *BlogArticleRepo) GetTotal(status int, keyword string, total *int64) (err error) {

	sql := "SELECT count(1) FROM blog_article a " +
		"JOIN blog_category c ON a.category_id = c.category_id " +
		"WHERE 1=1 "
	var where []interface{}
	if status != 0 {
		sql += " AND a.status = ? "
		where = append(where, status)
	}
	if keyword != "" {
		sql += " AND title like '?'"
		where = append(where, "%"+keyword+"%")
	}
	return r.db.Raw(sql, where...).Count(total).Error
}

type BlogArticle struct {
	ArticleId    int       `json:"article_id" gorm:"primary_key;" description:"主键"`
	Title        string    `json:"title" description:"文章标题"`
	Thumbnail    string    `json:"thumbnail" description:"缩略图"`
	Describe     string    `json:"describe" description:"描述"`
	Tags         string    `json:"tags" description:"关键词"`
	CategoryId   int       `json:"category_id" description:"关联栏目ID"`
	Content      string    `json:"content" description:"内容"`
	Status       int       `json:"status" description:"状态 0为草稿，1为显示，2为推荐，-1为删除(逻辑)"`
	CommentTotal int       `json:"comment_total" description:"评论数量"`
	AccessTotal  int       `json:"access_total" description:"访问量"`
	Url          string    `json:"url" description:"非原创的转载地址"`
	IsOriginal   int       `json:"is_original" description:"是否原创，0为不是，1为是"`
	Author       string    `json:"author" description:"作者"`
	Address      string    `json:"address" description:"地址"`
	AddDate      time.Time `json:"add_date" gorm:"add_date" description:"添加时间"`
	ChangeDate   time.Time `json:"change_date" gorm:"change_date" description:"修改时间"`
	IsDelete     int       `json:"is_delete" description:"是否删除"`
}

type BlogArticleListItem struct {
	BlogArticle
	CategoryName string `json:"category_name" gorm:"category_name" description:"栏目名称"`
}

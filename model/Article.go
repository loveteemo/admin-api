package model

import "time"

type Article struct {
	ArticleId  int       `json:"article_id" gorm:"primary_key;" description:"主键"`
	Title      string    `json:"title" description:"文章标题"`
	Thumbnail  string    `json:"thumbnail" description:"缩略图"`
	Describe   string    `json:"describe" description:"描述"`
	Tags       string    `json:"tags" description:"关键词"`
	CategoryId int       `json:"category_id" description:"关联栏目ID"`
	AddDate    time.Time `json:"add_date" gorm:"-" description:"添加时间"`
	Content    string    `json:"content" description:"内容"`
	Status     int       `json:"status" description:"状态 0为草稿，1为显示，2为推荐，-1为删除(逻辑)"`
	Author     string    `json:"author" description:"作者"`
	LikeNum    int       `json:"like_num" description:"点赞数量"`
	Hit        int       `json:"hit" description:"点击量"`
	Url        string    `json:"url" description:"非原创的转载地址"`
	IsOriginal int       `json:"is_original" description:"是否原创，0为不是，1为是"`
	ChangeDate time.Time `json:"change_date" gorm:"-" description:"修改时间"`
}

func (Article) TableName() string {
	return "blog_article"
}

func GetArticleList(page int, status int, keyword string, list *[]Article) (err error) {
	db = db.Model(&Article{})
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	if keyword != "" {
		db = db.Where("title like ?", "%"+keyword+"%")
	}

	return db.Offset((page - 1) * 10).Limit(10).Find(list).Error
}

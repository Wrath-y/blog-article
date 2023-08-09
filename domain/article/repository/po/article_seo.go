package po

import "time"

type ArticleSEO struct {
	Id          int64     `json:"id"`
	ArticleId   int64     `json:"article_id"`
	Title       string    `json:"title"`
	Keywords    string    `json:"keywords"`
	Description string    `json:"description"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}

func (*ArticleSEO) TableName() string {
	return "article_seo"
}

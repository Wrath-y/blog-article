package po

import "time"

type ArticleSEO struct {
	Id         int64     `json:"id"`
	ArticleId  int64     `json:"article_id"`
	Name       string    `json:"name"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

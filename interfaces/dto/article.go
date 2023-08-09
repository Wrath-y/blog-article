package dto

import "time"

type H struct{}

type ArticlesItem struct {
	Id           int64     `json:"id"`
	Title        string    `json:"title"`
	Image        string    `json:"image"`
	Intro        string    `json:"intro"`
	Hits         int       `json:"hits"`
	Source       int       `json:"source"`
	Tags         string    `json:"tags"`
	CommentCount int       `json:"comment_count"`
	CreateTime   time.Time `json:"create_time"`
}

type Article struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	Image        string `json:"image"`
	Html         string `json:"html"`
	Hits         int    `json:"hits"`
	Source       int    `json:"source"`
	Tags         string `json:"tags"`
	ArticleSEO   `json:"article_seo"`
	CommentCount int       `json:"comment_count"`
	CreateTime   time.Time `json:"create_time"`
}

type ArticleSEO struct {
	Title       string `json:"title"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
}

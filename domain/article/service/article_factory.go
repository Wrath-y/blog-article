package service

import (
	"article/domain/article/entity"
	"article/domain/article/repository/po"
)

type ArticleFactory struct {
}

func NewArticleFactory() ArticleFactory {
	return ArticleFactory{}
}

func (*ArticleFactory) CreateArticlePO(article entity.Article) po.Article {
	return po.Article{
		Id:         article.Id,
		Title:      article.Title,
		CreateTime: article.CreateTime,
		UpdateTime: article.UpdateTime,
	}
}

func (*ArticleFactory) CreateArticle(p po.Article) entity.Article {
	return entity.Article{
		Id:         p.Id,
		Title:      p.Title,
		Status:     p.Status,
		CreateTime: p.CreateTime,
		UpdateTime: p.UpdateTime,
	}
}

func (*ArticleFactory) CreateArticleSEOPO(seo entity.ArticleSEO) po.ArticleSEO {
	return po.ArticleSEO{
		Id:         seo.Id,
		ArticleId:  seo.ArticleId,
		Name:       seo.Name,
		Content:    seo.Content,
		CreateTime: seo.CreateTime,
		UpdateTime: seo.UpdateTime,
	}
}

func (*ArticleFactory) CreateArticleSEO(p po.ArticleSEO) entity.ArticleSEO {
	return entity.ArticleSEO{
		Id:         p.Id,
		ArticleId:  p.ArticleId,
		Name:       p.Name,
		Content:    p.Content,
		CreateTime: p.CreateTime,
		UpdateTime: p.UpdateTime,
	}
}

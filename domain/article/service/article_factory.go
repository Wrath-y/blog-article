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
		Image:      article.Image,
		Intro:      article.Intro,
		Html:       article.Html,
		Con:        article.Con,
		Hits:       article.Hits,
		Status:     article.Status,
		Source:     article.Source,
		Tags:       article.Tags,
		CreateTime: article.CreateTime,
		UpdateTime: article.UpdateTime,
	}
}

func (a *ArticleFactory) CreateArticleEntities(poList []*po.Article) []*entity.Article {
	res := make([]*entity.Article, 0, len(poList))
	for _, v := range poList {
		tmp := a.CreateArticleEntity(*v)
		res = append(res, &tmp)
	}

	return res
}

func (*ArticleFactory) CreateArticleEntity(article po.Article) entity.Article {
	return entity.Article{
		Id:         article.Id,
		Title:      article.Title,
		Image:      article.Image,
		Intro:      article.Intro,
		Html:       article.Html,
		Con:        article.Con,
		Hits:       article.Hits,
		Status:     article.Status,
		Source:     article.Source,
		Tags:       article.Tags,
		CreateTime: article.CreateTime,
		UpdateTime: article.UpdateTime,
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

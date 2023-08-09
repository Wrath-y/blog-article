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

//
//func (*ArticleFactory) CreateArticleSEOPO(seo valueobject.ArticleSEO) po.ArticleSEO {
//	return po.ArticleSEO{
//		Id:          seo.Id,
//		ArticleId:   seo.ArticleId,
//		Title:       seo.Title,
//		Keywords:    seo.Keywords,
//		Description: seo.Description,
//		CreateTime:  seo.CreateTime,
//		UpdateTime:  seo.UpdateTime,
//	}
//}
//
//func (*ArticleFactory) CreateArticleSEOEntity(p po.ArticleSEO) valueobject.ArticleSEO {
//	return valueobject.ArticleSEO{
//		Id:          p.Id,
//		ArticleId:   p.ArticleId,
//		Title:       p.Title,
//		Keywords:    p.Keywords,
//		Description: p.Description,
//		CreateTime:  p.CreateTime,
//		UpdateTime:  p.UpdateTime,
//	}
//}

package service

import (
	"article/domain/article/entity"
	"article/domain/article/service"
	"article/infrastructure/common/context"
)

type ArticleApplicationService struct {
	*context.Context
	service.ArticleDomainService
}

func NewArticleApplicationService(ctx *context.Context) *ArticleApplicationService {
	return &ArticleApplicationService{
		ctx,
		service.NewArticleDomainService(ctx),
	}
}

func (a *ArticleApplicationService) FindById(id int64) (entity.Article, error) {
	return a.ArticleDomainService.FindById(id)
}

func (a *ArticleApplicationService) Insert(article entity.Article) error {
	return a.ArticleDomainService.Insert(article)
}

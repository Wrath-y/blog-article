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

func (a *ArticleApplicationService) FindById(id int64, size int32) ([]*entity.Article, error) {
	return a.ArticleDomainService.FindById(id, size)
}

func (a *ArticleApplicationService) FindAll() ([]*entity.Article, error) {
	return a.ArticleDomainService.FindAll()
}

func (a *ArticleApplicationService) GetById(id int64) (entity.Article, error) {
	return a.ArticleDomainService.GetById(id)
}

package service

import (
	"article/domain/article/entity"
	"article/domain/article/service"
	"article/infrastructure/common/context"
)

type ArticleApplicationService struct {
	*context.Context
	articleDomainService service.ArticleDomainService
}

func NewArticleApplicationService(ctx *context.Context) *ArticleApplicationService {
	return &ArticleApplicationService{
		ctx,
		service.NewArticleDomainService(ctx),
	}
}

func (a *ArticleApplicationService) FindById(id int64, size int32) ([]*entity.Article, error) {
	return a.articleDomainService.FindById(id, size)
}

func (a *ArticleApplicationService) FindAll() ([]*entity.Article, error) {
	return a.articleDomainService.FindAll()
}

func (a *ArticleApplicationService) GetById(id int64) (entity.Article, error) {
	return a.articleDomainService.GetById(id)
}

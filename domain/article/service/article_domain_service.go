package service

import (
	"article/domain/article/entity"
	"article/domain/article/event"
	"article/domain/article/repository/facade"
	"article/domain/article/repository/persistence"
	"article/infrastructure/common/context"
	baseEvent "article/infrastructure/common/event"
)

type ArticleDomainService struct {
	*context.Context
	ArticleFactory
	facade.ArticleRepositoryI
	baseEvent.PublisherI
}

func NewArticleDomainService(ctx *context.Context) ArticleDomainService {
	return ArticleDomainService{
		Context:            ctx,
		ArticleFactory:     NewArticleFactory(),
		ArticleRepositoryI: persistence.NewArticleRepository(),
		PublisherI:         baseEvent.NewBasePublisher(),
	}
}

func (a *ArticleDomainService) FindById(id int64, size int32) ([]*entity.Article, error) {
	if size == 0 {
		size = 6
	}
	articles, err := a.ArticleRepositoryI.FindByLastId(id, size)
	if err != nil {
		a.Logger.ErrorL("获取文章列表失败", id, err.Error())
		return nil, err
	}

	return a.ArticleFactory.CreateArticleEntities(articles), nil
}

func (a *ArticleDomainService) FindAll() ([]*entity.Article, error) {
	articles, err := a.ArticleRepositoryI.FindAll()
	if err != nil {
		a.Logger.ErrorL("获取所有文章失败", "", err.Error())
		return nil, err
	}

	return a.ArticleFactory.CreateArticleEntities(articles), nil
}

func (a *ArticleDomainService) GetById(id int64) (entity.Article, error) {
	article, err := a.ArticleRepositoryI.GetById(id)
	if err != nil {
		a.Logger.ErrorL("获取文章详情失败", id, err.Error())
		return entity.Article{}, err
	}

	a.PublisherI = new(event.ArticleRead).Create("ArticleRead", id)
	if err := a.PublisherI.Publish(); err != nil {
		a.Logger.ErrorL("发布ArticleRead事件失败", id, err.Error())
	}

	return a.ArticleFactory.CreateArticleEntity(article), nil
}

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

func (a *ArticleDomainService) Insert(article entity.Article) error {
	_, err := a.ArticleRepositoryI.FindById(article.Id)
	if err != nil {
		return err
	}
	article.Create()
	return a.ArticleRepositoryI.Insert(a.ArticleFactory.CreateArticlePO(article))
}

func (a *ArticleDomainService) FindById(id int64) (entity.Article, error) {
	article, err := a.ArticleRepositoryI.FindById(id)
	if err != nil {
		a.Logger.ErrorL("获取文章详情失败", id, err.Error())
		return entity.Article{}, err
	}

	a.PublisherI = new(event.ArticleRead).Create("ArticleRead", id)
	if err := a.PublisherI.Publish(); err != nil {
		a.Logger.ErrorL("发布ArticleRead事件失败", id, err.Error())
	}

	return a.ArticleFactory.CreateArticle(article), nil
}

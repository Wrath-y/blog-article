package service

import (
	"article/domain/article/entity"
	"article/domain/article/event"
	"article/domain/article/repository/facade"
	"article/domain/article/repository/persistence"
	"article/domain/article/repository/po"
	"article/infrastructure/common/context"
	baseEvent "article/infrastructure/common/event"
	"github.com/go-redis/redis/v7"
)

type ArticleDomainService struct {
	*context.Context
	articleFactory     ArticleFactory
	articleCache       ArticleCache
	articleRepositoryI facade.ArticleRepositoryI
	publisherI         baseEvent.PublisherI
}

func NewArticleDomainService(ctx *context.Context) ArticleDomainService {
	return ArticleDomainService{
		Context:            ctx,
		articleFactory:     NewArticleFactory(),
		articleCache:       NewArticleCache(),
		articleRepositoryI: persistence.NewArticleRepository(),
		publisherI:         baseEvent.NewBasePublisher(),
	}
}

func (a *ArticleDomainService) FindById(id int64, size int32) ([]*entity.Article, error) {
	if size == 0 {
		size = 6
	}

	var err error
	articles := make([]*po.Article, 0)

	articles, err = a.articleCache.GetList(id, size)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if len(articles) > 0 {
		return a.articleFactory.CreateArticleEntities(articles), nil
	}

	articles, err = a.articleRepositoryI.FindByLastId(id, size)
	if err != nil {
		a.Logger.ErrorL("获取文章列表失败", id, err.Error())
		return nil, err
	}
	if err := a.articleCache.SetList(id, size, articles); err != nil {
		a.Logger.ErrorL("缓存文章列表失败", id, err.Error())
	}

	return a.articleFactory.CreateArticleEntities(articles), nil
}

func (a *ArticleDomainService) FindAll() ([]*entity.Article, error) {
	articles, err := a.articleRepositoryI.FindAll()
	if err != nil {
		a.Logger.ErrorL("获取所有文章失败", "", err.Error())
		return nil, err
	}

	return a.articleFactory.CreateArticleEntities(articles), nil
}

func (a *ArticleDomainService) GetById(id int64) (entity.Article, error) {
	defer func() {
		if err := a.publisherI.AddFunc(event.ArticleRead(
			func() error {
				return a.articleRepositoryI.HitsIncr(id)
			},
			func() error {
				return a.articleCache.HitsIncr(id)
			},
		)).Publish(a.Context); err != nil {
			a.Logger.ErrorL("发布ArticleRead事件失败", id, err.Error())
		}
	}()

	var err error
	article := po.Article{}

	article, err = a.articleCache.GetDetail(id)
	if err != nil && err != redis.Nil {
		a.Logger.ErrorL("获取文章详情缓存失败", id, err.Error())
		return entity.Article{}, err
	}
	if err == nil {
		return a.articleFactory.CreateArticleEntity(article), nil
	}

	article, err = a.articleRepositoryI.GetById(id)
	if err != nil {
		a.Logger.ErrorL("获取文章详情失败", id, err.Error())
		return entity.Article{}, err
	}

	article.Hits++
	if err := a.articleCache.SetDetail(id, article); err != nil {
		a.Logger.ErrorL("缓存文章详情失败", id, err.Error())
	}

	return a.articleFactory.CreateArticleEntity(article), nil
}

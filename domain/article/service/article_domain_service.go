package service

import (
	"article/domain/article/entity"
	"article/domain/article/event"
	"article/domain/article/repository/facade"
	"article/domain/article/repository/persistence"
	"article/domain/article/repository/po"
	"article/infrastructure/common/context"
	baseEvent "article/infrastructure/common/event"
	"article/infrastructure/util/consul"
	"article/infrastructure/util/grpcclient"
	"article/interfaces/proto"
	ctx "context"
	"errors"
	"github.com/go-redis/redis/v7"
	"time"
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
	if err != nil && !errors.Is(err, redis.Nil) {
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
			func() error {
				return a.articleCache.ClearList()
			},
		)).Publish(a.Context); err != nil {
			a.Logger.ErrorL("发布ArticleRead事件失败", id, err.Error())
		}
	}()

	var err error
	articleEntity := entity.Article{}
	articlePo := po.Article{}

	articleEntity, err = a.articleCache.GetDetail(id)
	if err != nil && !errors.Is(err, redis.Nil) {
		a.Logger.ErrorL("获取文章详情缓存失败", id, err.Error())
		return entity.Article{}, err
	}
	if err == nil {
		return articleEntity, nil
	}

	articlePo, err = a.articleRepositoryI.GetById(id)
	if err != nil {
		a.Logger.ErrorL("获取文章详情失败", id, err.Error())
		return entity.Article{}, err
	}

	articleEntity = a.articleFactory.CreateArticleEntity(articlePo)
	articleEntity.IncrHits()

	articleSEO, err := a.articleRepositoryI.GetSEOByArticleId(id)
	if err != nil {
		a.Logger.ErrorL("获取文章seo失败", id, err.Error())
		return articleEntity, nil
	}
	articleEntity.CreateArticleSEO(articleSEO)

	c, grpcClient, closeFunc, err := a.getCommentClient("comment")
	if err != nil {
		a.Logger.ErrorL("获取评论服务client失败", "", err.Error())
		return articleEntity, nil
	}
	defer closeFunc()

	req := &proto.OnlyArticleIdReq{
		ArticleId: id,
	}
	rpcResp, err := grpcClient.GetCountByArticleId(c, req)
	if err != nil {
		a.Logger.ErrorL("获取评论数量失败", "", err.Error())
	}
	if rpcResp != nil && rpcResp.Data != "" {
		articleEntity.CreateCommentCount(rpcResp.Data)
	}

	if err := a.articleCache.SetDetail(id, articleEntity); err != nil {
		a.Logger.ErrorL("缓存文章详情失败", id, err.Error())
	}

	return articleEntity, nil
}

func (a *ArticleDomainService) getCommentClient(serviceName string) (ctx.Context, proto.CommentClient, func(), error) {
	instance, err := consul.Client.GetHealthRandomInstance(serviceName)
	if err != nil {
		return nil, nil, nil, err
	}

	conn, err := grpcclient.NewClient(instance).GetHealthConn()
	if err != nil {
		return nil, nil, nil, err
	}

	connCloseFunc := func() {
		if err := conn.Close(); err != nil {
			a.Logger.ErrorL("grpc链接关闭失败", "", err.Error())
		}
	}

	grpcClient := proto.NewCommentClient(conn)

	c, cancel := ctx.WithTimeout(ctx.Background(), time.Second*3)
	closeFunc := func() {
		connCloseFunc()
		cancel()
	}

	return c, grpcClient, closeFunc, nil
}

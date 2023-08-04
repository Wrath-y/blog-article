package facade

import "article/domain/article/repository/po"

type ArticleRepositoryI interface {
	FindByLastId(lastId int64, limit int32) ([]*po.Article, error)
	FindAll() ([]*po.Article, error)
	GetById(id int64) (po.Article, error)
	GetHitsById(id int64) (po.Article, error)
	HitsIncr(id int64) error
}

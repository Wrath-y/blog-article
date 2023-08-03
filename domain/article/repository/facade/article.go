package facade

import "article/domain/article/repository/po"

type ArticleRepositoryI interface {
	Insert(po.Article) error
	Update(po.Article) error
	FindById(int64) (po.Article, error)
}

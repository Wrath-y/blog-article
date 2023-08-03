package persistence

import (
	"article/domain/article/repository/facade"
	"article/domain/article/repository/po"
	"article/infrastructure/util/db"
)

type ArticleRepository struct{}

func NewArticleRepository() facade.ArticleRepositoryI {
	return &ArticleRepository{}
}

func (*ArticleRepository) Insert(article po.Article) error {
	return db.Orm.Save(&article).Error
}

func (*ArticleRepository) Update(article po.Article) error {
	return db.Orm.Save(&article).Error
}

func (*ArticleRepository) FindById(id int64) (po.Article, error) {
	var articlePO po.Article
	return articlePO, db.Orm.Find(&articlePO, id).Error
}

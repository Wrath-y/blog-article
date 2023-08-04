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

func (*ArticleRepository) FindByLastId(lastId int64, limit int32) ([]*po.Article, error) {
	var article []*po.Article
	if lastId == 0 {
		return article, db.Orm.Raw("select * from article where status = 1 order by id desc limit ?", limit).Find(&article).Error
	}
	if lastId == -1 {
		return article, db.Orm.Raw("select * from article where status = 1 order by id desc").Find(&article).Error
	}
	return article, db.Orm.Raw("select * from article where id < ? and status = 1 order by id desc limit ?", lastId, limit).Find(&article).Error
}

func (*ArticleRepository) FindAll() ([]*po.Article, error) {
	var article []*po.Article
	return article, db.Orm.Raw("select * from article order by id desc").Find(&article).Error
}

func (*ArticleRepository) GetById(id int64) (po.Article, error) {
	article := po.Article{}
	return article, db.Orm.First(&article, id).Error
}

func (*ArticleRepository) GetHitsById(id int64) (po.Article, error) {
	article := po.Article{}
	return article, db.Orm.Raw("select id, hits from article where id = ?", id).First(&article).Error
}

func (*ArticleRepository) HitsIncr(id int64) error {
	return db.Orm.Exec("update article set hits = hits + 1 where id = ?", id).Error
}

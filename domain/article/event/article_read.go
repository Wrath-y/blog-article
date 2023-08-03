package event

import (
	"article/infrastructure/common/event"
	"strconv"
	"time"
)

type ArticleRead struct {
	event.Base
}

func (a *ArticleRead) Create(source string, articleId int64) *ArticleRead {
	a.Source = source
	a.Data = strconv.FormatInt(articleId, 10)
	a.CreateTime = time.Now()

	return a
}

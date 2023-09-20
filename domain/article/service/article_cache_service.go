package service

import (
	"article/domain/article/entity"
	"article/domain/article/repository/po"
	"article/infrastructure/util/def"
	"article/infrastructure/util/goredis"
	"article/infrastructure/util/util/highperf"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/go-redis/redis/v7"
	"strconv"
	"time"
)

const (
	ListStrKey    = "blog:article:list:%d:%d"
	DetailHashKey = "blog:article:%d"
	//SEODetailHashKey = "blog:article:seo:%d"
)

type ArticleCache struct {
}

func NewArticleCache() ArticleCache {
	return ArticleCache{}
}

func (*ArticleCache) GetList(id int64, size int32) ([]*po.Article, error) {
	articles := make([]*po.Article, 0)
	b, err := goredis.Client.Get(fmt.Sprintf(ListStrKey, id, size)).Bytes()
	if err != nil {
		return nil, err
	}

	if err := sonic.Unmarshal(b, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}

func (*ArticleCache) SetList(id int64, size int32, list []*po.Article) error {
	b, err := sonic.Marshal(list)
	if err != nil {
		return err
	}
	return goredis.Client.Set(fmt.Sprintf(ListStrKey, id, size), highperf.Bytes2str(b), time.Hour*24*7).Err()
}

func (*ArticleCache) GetDetail(id int64) (entity.Article, error) {
	article := new(entity.Article)
	m, err := goredis.Client.HGetAll(fmt.Sprintf(DetailHashKey, id)).Result()
	if err != nil {
		return *article, err
	}
	if len(m) == 0 {
		return *article, redis.Nil
	}

	article.Id, err = strconv.ParseInt(m["id"], 10, 64)
	article.Title = m["title"]
	article.Image = m["image"]
	article.Intro = m["intro"]
	article.Html = m["html"]
	article.Con = m["con"]
	article.Hits, err = strconv.Atoi(m["hits"])
	article.Tags = m["tags"]
	status, err := strconv.Atoi(m["status"])
	article.Status = int8(status)
	article.Source, err = strconv.Atoi(m["source"])
	article.CommentCount, err = strconv.Atoi(m["comment_count"])
	article.CreateTime, err = time.Parse(def.ISO8601Layout, m["create_time"])
	article.UpdateTime, err = time.Parse(def.ISO8601Layout, m["update_time"])

	if err := sonic.Unmarshal(highperf.Str2bytes(m["article_seo"]), &article.ArticleSEO); err != nil {
		return *article, err
	}

	return *article, nil
}

func (*ArticleCache) SetDetail(id int64, detail entity.Article) error {
	b, err := sonic.Marshal(detail)
	if err != nil {
		return err
	}
	m := make(map[string]any)
	if err = sonic.Unmarshal(b, &m); err != nil {
		return err
	}
	values := make([]any, 0, len(m)*2)
	for k, v := range m {
		if k == "article_seo" {
			v, err = sonic.Marshal(v)
			if err != nil {
				return err
			}

		}
		values = append(values, k, v)
	}
	key := fmt.Sprintf(DetailHashKey, id)
	if err := goredis.Client.HSet(key, values...).Err(); err != nil {
		return err
	}

	return goredis.Client.Expire(key, time.Hour*24*7).Err()
}

func (*ArticleCache) HitsIncr(id int64) error {
	key := fmt.Sprintf(DetailHashKey, id)
	if exists, err := goredis.Client.Exists(key).Result(); err != nil || exists == 0 {
		return err
	}
	if err := goredis.Client.HIncrBy(key, "hits", 1).Err(); err != nil {
		return err
	}

	return goredis.Client.Expire(key, time.Hour*24*7).Err()
}

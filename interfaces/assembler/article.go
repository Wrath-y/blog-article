package assembler

import (
	"article/domain/article/entity"
	"article/interfaces/dto"
	"article/interfaces/proto"
	"time"
)

func ToArticleEntity(articleDTO dto.Article) entity.Article {
	return entity.Article{
		Id:    articleDTO.Id,
		Title: articleDTO.Title,
	}
}

func ToArticleDTO(article entity.Article) dto.Article {
	return dto.Article{
		Id:    article.Id,
		Title: article.Title,
	}
}
func PbToArticleEntity(article *proto.ArticleBaseInfo) entity.Article {
	return entity.Article{
		Title:      article.Title,
		Image:      "",
		Intro:      "",
		Html:       "",
		Con:        "",
		Hits:       0,
		Status:     0,
		Source:     0,
		Tags:       "",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}

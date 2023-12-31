package assembler

import (
	"article/domain/article/entity"
	"article/interfaces/dto"
)

func ToArticleEntity(articleDTO dto.ArticlesItem) entity.Article {
	return entity.Article{
		Id:    articleDTO.Id,
		Title: articleDTO.Title,
	}
}

func ToArticleDTOs(articles []*entity.Article) []*dto.ArticlesItem {
	res := make([]*dto.ArticlesItem, 0, len(articles))
	for _, v := range articles {
		tmp := ToArticleListTmpDTO(*v)
		res = append(res, &tmp)
	}

	return res
}

func ToArticleListTmpDTO(article entity.Article) dto.ArticlesItem {
	return dto.ArticlesItem{
		Id:           article.Id,
		Title:        article.Title,
		Image:        article.Image,
		Intro:        article.Intro,
		Hits:         article.Hits,
		Source:       article.Source,
		Tags:         article.Tags,
		CommentCount: article.CommentCount,
		CreateTime:   article.CreateTime,
	}
}

func ToArticleDTO(article entity.Article) dto.Article {
	return dto.Article{
		Id:         article.Id,
		Title:      article.Title,
		Image:      article.Image,
		Html:       article.Html,
		Hits:       article.Hits,
		Source:     article.Source,
		Tags:       article.Tags,
		CreateTime: article.CreateTime,
		ArticleSEO: dto.ArticleSEO{
			Title:       article.ArticleSEO.Title,
			Keywords:    article.ArticleSEO.Keywords,
			Description: article.ArticleSEO.Description,
		},
		CommentCount: article.CommentCount,
	}
}

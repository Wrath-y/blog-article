package facade

import (
	"article/application/service"
	grpcCtx "article/infrastructure/common/context"
	"article/infrastructure/util/errcode"
	"article/interfaces/assembler"
	"article/interfaces/dto"
	"article/interfaces/proto"
	"article/launch/grpc/resp"
	"context"
)

type Article struct{}

func (*Article) FindById(ctx context.Context, req *proto.ArticleBaseInfo) (*proto.Response, error) {
	res, err := service.NewArticleApplicationService(grpcCtx.GetContext(ctx)).FindById(req.Id)
	if err != nil {
		return resp.FailWithErrCode(errcode.ArticleNotExists)
	}

	return resp.Success(assembler.ToArticleDTO(res))
}

func (*Article) Insert(ctx context.Context, req *proto.ArticleBaseInfo) (*proto.Response, error) {
	if err := service.NewArticleApplicationService(grpcCtx.GetContext(ctx)).Insert(assembler.PbToArticleEntity(req)); err != nil {
		return resp.FailWithErrCode(errcode.ArticleInsertFailed)
	}

	return resp.Success(dto.H{})
}

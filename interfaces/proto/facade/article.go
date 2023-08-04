package facade

import (
	"article/application/service"
	grpcCtx "article/infrastructure/common/context"
	"article/infrastructure/util/errcode"
	"article/interfaces/assembler"
	"article/interfaces/proto"
	"article/launch/grpc/resp"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Article struct{}

func (*Article) FindById(ctx context.Context, req *proto.FindByIdReq) (*proto.Response, error) {
	res, err := service.NewArticleApplicationService(grpcCtx.GetContext(ctx)).FindById(req.Id, req.Size)
	if err != nil {
		return resp.FailWithErrCode(errcode.ArticleFindFailed)
	}

	return resp.Success(assembler.ToArticleDTOs(res))
}

func (*Article) FindAll(ctx context.Context, _ *emptypb.Empty) (*proto.Response, error) {
	res, err := service.NewArticleApplicationService(grpcCtx.GetContext(ctx)).FindAll()
	if err != nil {
		return resp.FailWithErrCode(errcode.ArticleFindFailed)
	}
	return resp.Success(assembler.ToArticleDTOs(res))
}

func (*Article) GetById(ctx context.Context, req *proto.GetByIdReq) (*proto.Response, error) {
	res, err := service.NewArticleApplicationService(grpcCtx.GetContext(ctx)).GetById(req.Id)
	if err != nil {
		return resp.FailWithErrCode(errcode.ArticleNotExists)
	}
	return resp.Success(assembler.ToArticleDTO(res))
}

package middleware

import (
	grpcCtx "article/infrastructure/common/context"
	"context"
	"google.golang.org/grpc"
)

func UnaryContext() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		c := grpcCtx.NewContext(ctx)

		return handler(c, req)
	}
}

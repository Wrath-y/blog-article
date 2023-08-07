package facade

import (
	"article/infrastructure/util/consul"
	"article/interfaces/proto"
	"context"
	"github.com/spf13/viper"
	"log"
	"testing"
	"time"
)

func TestFindAll(t *testing.T) {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../../")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	consul.Setup()

	conn, err := consul.Client.GetGRPCHealthConn("article")
	defer conn.Close()

	var grpcClient proto.ArticleClient
	grpcClient = proto.NewArticleClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	r, err := grpcClient.FindById(ctx, &proto.FindByIdReq{
		Id:   0,
		Size: 1,
	})
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(r)
}

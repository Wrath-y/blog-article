package grpc

import (
	"article/infrastructure/util/consul"
	defgrpc "article/infrastructure/util/def/grpc"
	"article/infrastructure/util/logging"
	"article/interfaces/proto"
	"article/interfaces/proto/facade"
	middleware2 "article/launch/grpc/middleware"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func RunGrpc() {
	lis, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		// options
		grpc.InitialWindowSize(defgrpc.InitialWindowSize),
		grpc.InitialConnWindowSize(defgrpc.InitialConnWindowSize),
		grpc.MaxSendMsgSize(defgrpc.MaxSendMsgSize),
		grpc.MaxRecvMsgSize(defgrpc.MaxRecvMsgSize),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			PermitWithoutStream: true,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    defgrpc.KeepAliveTime,
			Timeout: defgrpc.KeepAliveTimeout,
		}),
		// middlewares
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			middleware2.UnaryRecover(),
			middleware2.UnaryContext(),
			middleware2.UnaryLogger(),
		)))

	// 在gRPC服务器注册我们的服务
	grpc_health_v1.RegisterHealthServer(grpcServer, facade.NewHealthCheckService())
	proto.RegisterArticleServer(grpcServer, &facade.Article{})

	go func() {
		//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("grpcServer.Serve err: %v", err)
		}
		log.Println("Shutdown grpcServer.Serve")
	}()

	registerService()

	logging.New().Info("Has Start", "", viper.GetString("app.env"))

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	globalDestroy()

	grpcServer.GracefulStop()
}

func registerService() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal("获取IP地址失败:", err.Error())
	}
	ip := ""
	// 排除回环地址和IPv6地址，只输出本机的IPv4地址
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			log.Println("本机IP地址:", ipnet.IP.String())
			ip = ipnet.IP.String()
			break
		}
	}
	if ip == "" {
		log.Fatal("IP地址未获取到:", err.Error())
	}
	serviceInstance := consul.NewServiceInstance(strconv.FormatInt(time.Now().Unix(), 10), "article", "grpc", ip, 8000, false, map[string]string{})
	if err := consul.Client.Register(serviceInstance); err != nil {
		log.Fatalf("consul.Register err: %v", err)
	}
}

func globalDestroy() {
	if err := consul.Client.Deregister(); err != nil {
		logging.New().ErrorL("consul deregister failed", "", err.Error())
	}
}

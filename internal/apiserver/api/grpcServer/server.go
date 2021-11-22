/*
   @Author:huolun
   @Date:2021/11/22
   @Description
*/
package grpcServer

import (
	"context"
	"fmt"
	"git.cai-inc.com/devops/someProject/internal/apiserver/api/grpcServer/protos/demo"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	*grpc.Server
}

func (*Server) Ping(ctx context.Context, req *demo.Request) (*demo.Response, error) {
	return &demo.Response{Value: "pong"}, nil
}

func NewGrpcServer() *Server {
	fmt.Println("starting gRPC server...")
	grpcServer := grpc.NewServer()
	demo.RegisterDemoServerServer(grpcServer, &Server{})
	return &Server{
		grpcServer,
	}

}
func (s *Server)Run(Addr string) {
	lis, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}
	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		log.Println(fmt.Sprintf("start grpc server on %s",Addr))
		err = s.Server.Serve(lis)
		if err!= nil{
			log.Fatalf("failed to serve: %v \n", err)
		}
		return err
	})
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	s.GracefulStop()
	log.Println("grpc exiting")

}

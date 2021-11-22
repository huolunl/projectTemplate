/*
   @Author:huolun
   @Date:2021/11/22
   @Description
*/
package grpcServer

import (
	"context"
	"git.cai-inc.com/devops/someProject/internal/apiserver/api/grpcServer/protos/demo"
	"google.golang.org/grpc"
	"log"
	"time"
)

func NewGrpcDemoClient(Addr string) {
	conn, err := grpc.Dial(Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	client := demo.NewDemoServerClient(conn)
	ctx,cancel := context.WithTimeout(context.Background(),time.Second*5)
	defer cancel()
	resp,err:=client.Ping(ctx,&demo.Request{Id: "1"})
	if err!= nil{
		log.Println(err)
	}
	log.Println(resp)

}

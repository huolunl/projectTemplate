/*
   @Author:huolun
   @Date:2021/11/22
   @Description
*/
package main

import "git.cai-inc.com/devops/someProject/internal/apiserver/api/grpcServer"

func main()  {
	grpcServer.NewGrpcServer().Run("localhost:50051")
}

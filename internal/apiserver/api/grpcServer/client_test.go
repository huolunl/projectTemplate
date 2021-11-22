/*
   @Author:huolun
   @Date:2021/11/22
   @Description
*/
package grpcServer

import "testing"

func TestNewGrpcDemoClient(t *testing.T) {
	NewGrpcDemoClient("localhost:50051")
}

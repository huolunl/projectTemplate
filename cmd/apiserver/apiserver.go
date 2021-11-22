package main

import "git.cai-inc.com/devops/someProject/internal/apiserver/api/genericAPIserver"

func main() {
	genericAPIserver.NewHttpServer(":8080").Run()
}

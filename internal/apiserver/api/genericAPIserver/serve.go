/*
   @Author:huolun
   @Date:2021/11/22
   @Description
*/
package genericAPIserver

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)
type HttpServer struct {
	mu sync.Mutex
	*http.Server
}

func NewHttpServer(Addr string) *HttpServer {
	router := gin.Default()
	RouterRegister(router)
	return &HttpServer{
		sync.Mutex{},
		&http.Server{
			Addr:    Addr,
			Handler: router,
		},
	}
}

func (h *HttpServer)Run() {
	go func() {
		// service connections
		if err := h.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := h.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}
